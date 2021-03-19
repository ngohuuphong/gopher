package main

import (
	"fmt"
	"golang-test-1/models"
	"golang-test-1/routes"
	"golang-test-1/utils"
	"log"
	"net/http"
	"os"
)

func main() {
	models.TestConnection()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Listening Port %s\n", port)
	utils.LoadTemplates("views/*.html")
	r := routes.NewRouter()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
