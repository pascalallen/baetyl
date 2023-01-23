package Security

import "github.com/pascalallen/baetyl/src/Adapter/Security"

type TokenEncoder interface {
	// TODO: Rely on a more abstract claims type
	Encode(claims Security.UserClaims) (string, error)
}
