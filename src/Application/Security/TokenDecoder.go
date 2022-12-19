package Security

import "github.com/pascalallen/Baetyl/src/Adapter/Security"

type TokenDecoder interface {
	// TODO: Rely on a more abstract claims type
	Decode(tokenString string) (*Security.UserClaims, error)
}
