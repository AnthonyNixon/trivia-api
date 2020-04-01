package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var psqlInfo string
var db *sql.DB

func Initialize() {
	log.Print("Initializing Database...")
	var DB_USER = os.Getenv("TRIVIA_DB_USER")
	var DB_PASS = os.Getenv("TRIVIA_DB_PASS")
	var DB_HOST = os.Getenv("TRIVIA_DB_HOST")
	var DB_NAME = os.Getenv("TRIVIA_DB_NAME")

	error := false
	if DB_USER == "" {
		log.Printf("TRIVIA_DB_USER environment variable not set")
		error = true
	}
	if DB_PASS == "" {
		log.Printf("TRIVIA_DB_PASS environment variable not set")
		error = true
	}
	if DB_NAME == "" {
		log.Printf("TRIVIA_DB_NAME environment variable not set")
		error = true
	}
	if DB_HOST == "" {
		log.Printf("TRIVIA_DB_HOST environment variable not set")
		error = true
	}

	if error {
		log.Fatal("Missing Database login information. Not Starting.")
	}

	setPSQLInfo(DB_HOST, DB_USER, DB_PASS, DB_NAME)

	dbInstance, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Printf("failed to open DB connection: %v", err.Error())
	}

	SetConnection(dbInstance)

	available, err := IsConnectionAvailable(5)
	if !available {
		log.Fatalf("Failed to Ping DB connection: %v", err.Error())
	}

	log.Print("done")
}

func IsConnectionAvailable(retries int) (bool, error) {
	err := Connection().Ping()
	if err != nil {
		if retries == 0 {
			return false, err
		} else {
			time.Sleep(2 * time.Second)
			return IsConnectionAvailable(retries - 1)
		}
	}

	return true, nil
}

func setPSQLInfo(host string, user string, pass string, dbname string) {
	psqlInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, "5432", user, pass, dbname)
}

func SetConnection(dbInstance *sql.DB) {
	db = dbInstance
}

func Connection() (connection *sql.DB) {
	return db
}
