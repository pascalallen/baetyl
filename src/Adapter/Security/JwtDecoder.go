package Security

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

type JwtDecoder struct {
	SecretKey string
}

func (decoder JwtDecoder) Decode(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(decoder.SecretKey), nil
	})

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("error parsing JWT: %v", err)
}
