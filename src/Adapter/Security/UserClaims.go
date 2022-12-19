package Security

import "github.com/golang-jwt/jwt/v4"

type UserClaims struct {
	Id    string `json:"id"`
	First string `json:"first"`
	Last  string `json:"last"`
	Iat   int    `json:"iat"`
	jwt.RegisteredClaims
}
