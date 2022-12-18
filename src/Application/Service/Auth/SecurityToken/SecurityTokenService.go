package SecurityToken

import (
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/Crypto"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/SecurityToken"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/User"
)

type SecurityTokenService interface {
	GetById(id ulid.ULID) (*SecurityToken.SecurityToken, error)
	GetByCrypto(crypto Crypto.Crypto) (*SecurityToken.SecurityToken, error)
	GetAllForUser(user User.User) (*[]SecurityToken.SecurityToken, error)
	Add(securityToken *SecurityToken.SecurityToken) error
	Remove(securityToken *SecurityToken.SecurityToken) error
	UpdateOrAdd(securityToken *SecurityToken.SecurityToken) error
	ClearExpiredTokens() error
}
