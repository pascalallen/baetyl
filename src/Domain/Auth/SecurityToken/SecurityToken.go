package SecurityToken

import (
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/baetyl/src/Adapter/Database/Type"
	"github.com/pascalallen/baetyl/src/Domain/Auth/Crypto"
	"github.com/pascalallen/baetyl/src/Domain/Auth/SecurityTokenType"
	"github.com/pascalallen/baetyl/src/Domain/Auth/User"
	"time"
)

type SecurityToken struct {
	Id        Type.GormUlid                       `json:"id" gorm:"primaryKey;size:26;not null"`
	Crypto    Crypto.Crypto                       `json:"crypto" gorm:"size:64;not null"`
	Type      SecurityTokenType.SecurityTokenType `json:"type" gorm:"size:10;not null"`
	UserId    Type.GormUlid                       `json:"-"`
	User      User.User                           `json:"user" gorm:"foreignKey:UserId;references:Id;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time                           `json:"created_at" gorm:"not null"`
	ExpiresAt time.Time                           `json:"expires_at" gorm:"not null"`
}

func GenerateReset(id ulid.ULID, user User.User, expiresAt time.Time) *SecurityToken {
	crypto := Crypto.Generate()
	createdAt := time.Now()

	return &SecurityToken{
		Id:        Type.GormUlid(id),
		Crypto:    crypto,
		Type:      SecurityTokenType.RESET,
		User:      user,
		CreatedAt: createdAt,
		ExpiresAt: expiresAt,
	}
}

func GenerateRefresh(id ulid.ULID, user User.User, expiresAt time.Time) *SecurityToken {
	crypto := Crypto.Generate()
	createdAt := time.Now()

	return &SecurityToken{
		Id:        Type.GormUlid(id),
		Crypto:    crypto,
		Type:      SecurityTokenType.REFRESH,
		User:      user,
		CreatedAt: createdAt,
		ExpiresAt: expiresAt,
	}
}

func (s *SecurityToken) IsExpired() bool {
	return s.ExpiresAt.Equal(time.Now()) || s.ExpiresAt.Before(time.Now())
}
