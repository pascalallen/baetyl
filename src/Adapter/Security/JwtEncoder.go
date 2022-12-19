package Security

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

type JwtEncoder struct {
	SecretKey     string
	SigningMethod *jwt.SigningMethodHMAC
}

func (encoder JwtEncoder) Encode(claims UserClaims) (string, error) {
	token := jwt.NewWithClaims(encoder.SigningMethod, claims)

	signedString, err := token.SignedString([]byte(encoder.SecretKey))
	if err != nil {
		return "", fmt.Errorf("failed to create and sign JWT: %v", err)
	}

	return signedString, nil
}
