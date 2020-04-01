package session

import (
	"fmt"
	"net/http"
	"time"

	"github.com/AnthonyNixon/trivia-api/cmd/utils"

	"github.com/AnthonyNixon/trivia-api/cmd/services/database"

	"github.com/AnthonyNixon/trivia-api/cmd/utils/httperr"
)

type Interface interface {
	CreateNewSession(name string, description string, start time.Time) (session Session, error httperr.HttpErr)
}

type Service struct{}

const ServiceKey = "SessionService"

func (s Service) CreateNewSession(name string, description string, start time.Time) (session Session, error httperr.HttpErr) {
	code, codeErr := getUniqueSessionCode()
	if codeErr != nil {
		return session, httperr.New(http.StatusInternalServerError, fmt.Sprintf("Failed to generate new code: %s", codeErr.Error()))
	}

	err := database.Connection().QueryRow(INSERT_NEW_SESSION, code, name, description, start).Scan(&session.Id)
	if err != nil {
		return session, httperr.New(http.StatusInternalServerError, fmt.Sprintf("Failed to insert new session: %s", err.Error()))
	}

	session.Name = name
	session.Description = description
	session.Code = code
	session.StartTime = start

	return session, nil
}

func isSessionCodeUnique(code string) (isUnique bool, err error) {
	var count int
	err = database.Connection().QueryRow(COUNT_OF_CODE, code).Scan(&count)
	if err != nil {
		return isUnique, err
	}

	isUnique = count == 0
	return isUnique, err
}

func getUniqueSessionCode() (code string, err error) {
	code = utils.NewSessionCode()
	isUnique, err := isSessionCodeUnique(code)
	if err != nil {
		return code, err
	}

	for !isUnique {
		code = utils.NewSessionCode()
		isUnique, err = isSessionCodeUnique(code)
		if err != nil {
			return code, err
		}
	}

	return code, nil
}
