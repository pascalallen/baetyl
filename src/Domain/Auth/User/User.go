package User

import (
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/PasswordHash"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/Permission"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/Role"
	"time"
)

type User struct {
	Id           ulid.ULID                 `json:"id" gorm:"primaryKey;size:26;not null"`
	FirstName    string                    `json:"first_name" gorm:"size:100;not null"`
	LastName     string                    `json:"last_name" gorm:"size:100;not null"`
	EmailAddress string                    `json:"email_address" gorm:"size:100;not null"`
	PasswordHash PasswordHash.PasswordHash `json:"-" gorm:"column:password;size:255;default:null"`
	Roles        []Role.Role               `json:"roles,omitempty" gorm:"many2many:user_roles"`
	CreatedAt    time.Time                 `json:"created_at" gorm:"not null"`
	ModifiedAt   time.Time                 `json:"modified_at" gorm:"not null"`
	// TODO: Determine how to make nullable/optional
	DeletedAt time.Time `json:"deleted_at,omitempty" gorm:"default:null"`
}

func Register(firstName string, lastName string, emailAddress string) *User {
	id := ulid.Make()
	createdAt := time.Now()

	return &User{
		Id:           id,
		FirstName:    firstName,
		LastName:     lastName,
		EmailAddress: emailAddress,
		CreatedAt:    createdAt,
		ModifiedAt:   createdAt,
	}
}

func (u *User) UpdateFirstName(firstName string) {
	u.FirstName = firstName
	u.ModifiedAt = time.Now()
}

func (u *User) UpdateLastName(lastName string) {
	u.LastName = lastName
	u.ModifiedAt = time.Now()
}

func (u *User) UpdateEmailAddress(emailAddress string) {
	u.EmailAddress = emailAddress
	u.ModifiedAt = time.Now()
}

func (u *User) SetPasswordHash(passwordHash PasswordHash.PasswordHash) {
	u.PasswordHash = passwordHash
	u.ModifiedAt = time.Now()
}

func (u *User) AddRole(role Role.Role) {
	for _, r := range u.Roles {
		if r.Id == role.Id {
			return
		}
	}

	u.Roles = append(u.Roles, role)
	u.ModifiedAt = time.Now()
}

func (u *User) RemoveRole(role Role.Role) {
	for i, r := range u.Roles {
		if r.Id == role.Id {
			u.Roles[i] = u.Roles[len(u.Roles)-1]
		}
	}

	u.Roles = u.Roles[:len(u.Roles)-1]
}

func (u *User) HasRole(name string) bool {
	for _, r := range u.Roles {
		if r.Name == name {
			return true
		}
	}

	return false
}

func (u *User) Permissions() []Permission.Permission {
	var permissions []Permission.Permission
	for _, r := range u.Roles {
		permissions = append(permissions, r.Permissions...)
	}

	return permissions
}

func (u *User) HasPermission(name string) bool {
	for _, p := range u.Permissions() {
		if p.Name == name {
			return true
		}
	}

	return false
}

func (u *User) IsDeleted() bool {
	return !u.DeletedAt.IsZero()
}

func (u *User) Delete() {
	u.DeletedAt = time.Now()
	u.ModifiedAt = u.DeletedAt
}

func (u *User) Restore() {
	u.DeletedAt = time.Time{}
	u.ModifiedAt = time.Now()
}
