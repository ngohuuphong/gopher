package api

import (
	"fmt"
	"log"
	"net/http"

	"golang-test-4-fullstack-mysql-jwt/api/router"
	"golang-test-4-fullstack-mysql-jwt/auto"
	"golang-test-4-fullstack-mysql-jwt/config"
)

func init() {
	config.Load()
	auto.Load()
}

// Run message
func Run() {
	fmt.Printf("\n\tListening [::]:%d", config.PORT)
	listen(config.PORT)
}

func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router.LoadCORs(r)))
}
