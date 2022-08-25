package Role

import (
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/Permission"
	"time"
)

type Role struct {
	Id          ulid.ULID               `json:"id" gorm:"primaryKey;size:26;not null"`
	Name        string                  `json:"name"`
	Permissions []Permission.Permission `json:"permissions,omitempty" gorm:"many2many:role_permissions"`
	CreatedAt   time.Time               `json:"created_at"`
	ModifiedAt  time.Time               `json:"modified_at"`
}

func Define(id ulid.ULID, name string) *Role {
	createdAt := time.Now()

	return &Role{
		Id:         id,
		Name:       name,
		CreatedAt:  createdAt,
		ModifiedAt: createdAt,
	}
}

func (r *Role) UpdateName(name string) {
	r.Name = name
	r.ModifiedAt = time.Now()
}

func (r *Role) AddPermission(permission Permission.Permission) {
	for _, p := range r.Permissions {
		if p.Id == permission.Id {
			return
		}
	}

	r.Permissions = append(r.Permissions, permission)
	r.ModifiedAt = time.Now()
}

func (r *Role) RemovePermission(permission Permission.Permission) {
	for i, p := range r.Permissions {
		if p.Id == permission.Id {
			r.Permissions[i] = r.Permissions[len(r.Permissions)-1]
		}
	}

	r.Permissions = r.Permissions[:len(r.Permissions)-1]
}

func (r *Role) HasPermission(name string) bool {
	for _, p := range r.Permissions {
		if p.Name == name {
			return true
		}
	}

	return false
}
