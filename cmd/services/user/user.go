package user

import (
	"fmt"
	"github.com/AnthonyNixon/trivia-api/cmd/services/database"
	"github.com/AnthonyNixon/trivia-api/cmd/utils/httperr"
	"net/http"
)

type Interface interface {
	CreateNewUser(name string) (user User, error httperr.HttpErr)
}

type Service struct{}

const ServiceKey = "UserService"


func (s Service) CreateNewUser(name string) (user User, error httperr.HttpErr) {
	err := database.Connection().QueryRow(INSERT_NEW_USER, name).Scan(&user.Id)
	if err != nil {
		return user, httperr.New(http.StatusInternalServerError, fmt.Sprintf("Failed to insert new user: %s", err.Error()))
	}

	user.Name = name

	return user, nil
}