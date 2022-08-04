package User

import (
	"github.com/oklog/ulid/v2"
)

type UserRepository interface {
	GetById(id ulid.ULID) *User
	GetByEmailAddress(emailAddress string) *User
	// TODO: Include pagination
	GetAll(includeDeleted bool) []*User
	Add(user *User)
	Remove(user *User)
}
