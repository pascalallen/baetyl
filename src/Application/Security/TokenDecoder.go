package Security

import "github.com/pascalallen/baetyl/src/Adapter/Security"

type TokenDecoder interface {
	// TODO: Rely on a more abstract claims type
	Decode(tokenString string) (*Security.UserClaims, error)
}
