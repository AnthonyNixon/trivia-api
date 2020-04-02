package token

import (
	"github.com/AnthonyNixon/trivia-api/cmd/services/user"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/AnthonyNixon/trivia-api/cmd/utils/httperr"
	)

var JWT_SIGNING_KEY []byte

const TOKEN_VALID_TIME = 12 * time.Hour // TODO: Change to 5 minutes for prod

func Initialize() {
	log.Print("Initializing Authentication")
	signingKey := os.Getenv("JWT_SIGNING_KEY")
	if signingKey == "" {
		log.Fatal("No Signing Key Present.")
	}

	JWT_SIGNING_KEY = []byte(signingKey)
	log.Print("done")
}

type Interface interface {
	NewUserToken(user user.User) (token string, error httperr.HttpErr)
	GetUserIdFromToken(token string) (userId int, err httperr.HttpErr)
}

type Service struct{}

const ServiceKey = "TokenService"

func (s Service) NewUserToken(user user.User) (token string, error httperr.HttpErr) {
	var jwtKey = JWT_SIGNING_KEY

	expirationTime := time.Now().Add(TOKEN_VALID_TIME)
	claims := &TriviaClaims{
		UserId:      user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)
	if err != nil {
		return "", httperr.New(http.StatusInternalServerError, err.Error())
	}

	return token, nil
}

func (s Service) GetUserIdFromToken(token string) (userId int, retErr httperr.HttpErr) {
	claims := &TriviaClaims{}
	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return JWT_SIGNING_KEY, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return userId, httperr.New(http.StatusUnauthorized, "signature invalid, "+err.Error())
		}

		if !tkn.Valid {
			return userId, httperr.New(http.StatusUnauthorized, err.Error())
		}

		return userId, httperr.New(http.StatusBadRequest, "Invalid JWT token, "+err.Error())

	}

	return claims.UserId, nil
}
