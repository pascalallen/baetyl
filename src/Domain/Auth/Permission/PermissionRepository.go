package Permission

import "github.com/oklog/ulid/v2"

type PermissionRepository interface {
	GetById(id ulid.ULID) *Permission
	GetByName(name string) *Permission
	// TODO: Include pagination
	GetAll() []*Permission
	Add(permission *Permission)
	Remove(permission *Permission)
}
