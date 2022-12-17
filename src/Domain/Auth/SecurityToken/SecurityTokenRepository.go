package SecurityToken

import (
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/Crypto"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/User"
)

type SecurityTokenRepository interface {
	GetById(id ulid.ULID) (*SecurityToken, error)
	GetByCrypto(crypto Crypto.Crypto) (*SecurityToken, error)
	GetAllForUser(user User.User) (*[]SecurityToken, error)
	Add(securityToken *SecurityToken) error
	Remove(securityToken *SecurityToken) error
	UpdateOrAdd(securityToken *SecurityToken) error
	ClearExpiredTokens() error
}
