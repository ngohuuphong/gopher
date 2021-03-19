package controllers

import (
	"encoding/json"
	"fmt"
	"golang-test-2-api/models"
	"golang-test-2-api/utils"
	"golang-test-2-api/validations"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"errors"
)

var (
	ErrInvalidCash = errors.New("Transferred value is invalid")
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := models.GetTransactions()
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(w, transactions)
}

func PostTransaction(w http.ResponseWriter, r *http.Request) {
	transaction, err := verifyTransaction(r)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	_, err = models.NewTransaction(transaction)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, utils.DefaultResponse{"Transaction completed successfully!", http.StatusCreated})
}

func verifyTransaction(r *http.Request) (models.Transaction, error) {
	// receives the public key from the wallet that will receive the amount
	params := mux.Vars(r)
	targetKey := params["public_key"]
	target, err := models.GetWalletByPublicKey(targetKey)
	if err != nil {
		return models.Transaction{}, err
	}
	// receives the json from the wallet that will send an amount with the balance and the public key
	body, _ := ioutil.ReadAll(r.Body)
	var origin models.Wallet
	err = json.Unmarshal(body, &origin)
	if err != nil {
		return models.Transaction{}, err
	}
	// verification structure, whether the portfolio exists
	originVerify, err := models.GetWalletByPublicKey(origin.PublicKey)
	if err != nil {
		return models.Transaction{}, err
	}
	if validations.IsEmpty(target.PublicKey) || validations.IsEmpty(originVerify.PublicKey) {
		return models.Transaction{}, models.ErrWalletNotFound
	}
	// check if the balance to be transferred is greater than the balance in the wallet or less than zero
	if origin.Balance > originVerify.Balance || origin.Balance < 0 {
		return models.Transaction{}, ErrInvalidCash
	}
	var transaction models.Transaction
	transaction.Cash = origin.Balance
	transaction.Message = fmt.Sprintf("%s transferred %.2f $, for %s", originVerify.User.Nickname, origin.Balance, target.User.Nickname)
	transaction.Origin = origin
	transaction.Target = target
	return transaction, nil
}
