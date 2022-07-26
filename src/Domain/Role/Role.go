package Role

import (
	"github.com/google/uuid"
	"github.com/pascalallen/Baetyl/src/Domain/Permission"
	"time"
)

type Role struct {
	id          uuid.UUID               `validate:"required,uuid"`
	name        string                  `validate:"required"`
	permissions []Permission.Permission `validate:"required"`
	createdAt   time.Time               `validate:"required,datetime"`
	modifiedAt  time.Time               `validate:"required,datetime"`
}

func Define(name string) Role {
	id := uuid.New()
	createdAt := time.Now()

	return Role{
		id:          id,
		name:        name,
		permissions: nil,
		createdAt:   createdAt,
		modifiedAt:  createdAt,
	}
}

func (r Role) Id() uuid.UUID {
	return r.id
}

func (r Role) Name() string {
	return r.name
}

func (r Role) UpdateName(name string) {
	r.name = name
	r.modifiedAt = time.Now()
}

func (r Role) Permissions() []Permission.Permission {
	return r.permissions
}

func (r Role) AddPermission(permission Permission.Permission) {
	// TODO: Verify that this works
	for _, p := range r.permissions {
		if p.Id() == permission.Id() {
			return
		}
	}

	r.permissions = append(r.permissions, permission)
	r.modifiedAt = time.Now()
}

func (r Role) RemovePermission(permission Permission.Permission) {
	// TODO
}

func (r Role) HasPermission(name string) bool {
	for _, p := range r.permissions {
		if p.Name() == name {
			return true
		}
	}

	return false
}

func (r Role) CreatedAt() time.Time {
	return r.createdAt
}

func (r Role) ModifiedAt() time.Time {
	return r.modifiedAt
}
