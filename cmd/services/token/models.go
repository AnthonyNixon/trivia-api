package token

import "github.com/dgrijalva/jwt-go"

type TriviaClaims struct {
	UserId      int    `json:"userId"`
	jwt.StandardClaims
}
