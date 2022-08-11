package Role

import "github.com/oklog/ulid/v2"

type RoleRepository interface {
	GetById(id ulid.ULID) *Role
	GetByName(name string) *Role
	// TODO: Include pagination
	GetAll() []*Role
	Add(role *Role)
	Remove(role *Role)
}
