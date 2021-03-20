package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func BuildDSN() string {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "3307"
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		user = "root"
	}

	pass := os.Getenv("DB_PASSWORD")
	if pass == "" {
		pass = "@root"
	}

	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		dbname = "supertest"
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, pass, host, port, dbname)
}
