package User

import (
	"github.com/google/uuid"
)

type UserRepository interface {
	GetById(id uuid.UUID) *User
	GetByEmailAddress(emailAddress string) *User
	// TODO: Include pagination
	GetAll(includeDeleted bool) []*User
	Add(user *User)
	Remove(user *User)
}
