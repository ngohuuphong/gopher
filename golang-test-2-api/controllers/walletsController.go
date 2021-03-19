package controllers

import (
	"golang-test-2-api/models"
	"golang-test-2-api/utils"
	"io/ioutil"
	"net/http"

	"encoding/json"
	"strconv"

	"github.com/gorilla/mux"
)

func GetWallets(w http.ResponseWriter, r *http.Request) {
	wallets, err := models.GetWallets()
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(w, wallets)
}

func GetWallet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	wallet, err := models.GetWalletByPublicKey(params["public_key"])
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(w, wallet)
}

func PutWallet(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	add, _ := strconv.ParseBool(keys.Get("add"))
	params := mux.Vars(r)
	var wallet models.Wallet
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &wallet)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	wallet.PublicKey = params["public_key"]
	var rows int64
	if add {
		rows, err = models.AddBalance(wallet)
	} else {
		rows, err = models.UpdateWallet(wallet)
	}
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, rows)
}
