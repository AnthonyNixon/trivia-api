package main

import (
	"fmt"
	"github.com/AnthonyNixon/trivia-api/cmd/handlers/users"
	"log"
	"os"

	"github.com/AnthonyNixon/trivia-api/cmd/services/token"

	"github.com/AnthonyNixon/trivia-api/cmd/handlers/sessions"

	"github.com/AnthonyNixon/trivia-api/cmd/services/database"

	"github.com/AnthonyNixon/trivia-api/cmd/handlers/up"
	"github.com/AnthonyNixon/trivia-api/cmd/services/router"
)

var PORT = ""

func init() {
	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	database.Initialize()
	token.Initialize()
}

func main() {
	router := router.New()
	up.AddUpV1(router)
	sessions.AddSessionsV1(router)
	users.AddUsersV1(router)

	log.Printf("Running Trivia API on :%s...", PORT)

	err := router.Run(fmt.Sprintf(":%s", PORT)) // listen and serve on 0.0.0.0:8080
	if err != nil {
		log.Fatal(err.Error())
	}
}
