package User

import (
	"github.com/google/uuid"
	"github.com/pascalallen/Baetyl/src/Domain/PasswordHash"
	"github.com/pascalallen/Baetyl/src/Domain/Permission"
	"github.com/pascalallen/Baetyl/src/Domain/Role"
	"time"
)

type User struct {
	id           uuid.UUID                 `validate:"required,uuid"`
	firstName    string                    `validate:"required"`
	lastName     string                    `validate:"required"`
	emailAddress string                    `validate:"required,email"`
	passwordHash PasswordHash.PasswordHash `validate:"required"`
	roles        []Role.Role               `validate:"required"`
	createdAt    time.Time                 `validate:"required,datetime"`
	modifiedAt   time.Time                 `validate:"required,datetime"`
	deletedAt    time.Time                 `validate:"datetime"`
}

func Register(firstName string, lastName string, emailAddress string) *User {
	id := uuid.New()
	createdAt := time.Now()

	return &User{
		id:           id,
		firstName:    firstName,
		lastName:     lastName,
		emailAddress: emailAddress,
		createdAt:    createdAt,
		modifiedAt:   createdAt,
	}
}

func (u *User) Id() uuid.UUID {
	return u.id
}

func (u *User) FirstName() string {
	return u.firstName
}

func (u *User) UpdateFirstName(firstName string) {
	u.firstName = firstName
	u.modifiedAt = time.Now()
}

func (u *User) LastName() string {
	return u.lastName
}

func (u *User) UpdateLastName(lastName string) {
	u.lastName = lastName
	u.modifiedAt = time.Now()
}

func (u *User) EmailAddress() string {
	return u.emailAddress
}

func (u *User) UpdateEmailAddress(emailAddress string) {
	u.emailAddress = emailAddress
	u.modifiedAt = time.Now()
}

func (u *User) PasswordHash() PasswordHash.PasswordHash {
	return u.passwordHash
}

func (u *User) UpdatePasswordHash(passwordHash PasswordHash.PasswordHash) {
	u.passwordHash = passwordHash
	u.modifiedAt = time.Now()
}

func (u *User) Roles() []Role.Role {
	return u.roles
}

func (u *User) AddRole(role Role.Role) {
	for _, r := range u.roles {
		if r.Id() == role.Id() {
			return
		}
	}

	u.roles = append(u.roles, role)
	u.modifiedAt = time.Now()
}

func (u *User) RemoveRole(role Role.Role) {
	for i, r := range u.roles {
		if r.Id() == role.Id() {
			u.roles[i] = u.roles[len(u.roles)-1]
		}
	}

	u.roles = u.roles[:len(u.roles)-1]
}

func (u *User) HasRole(name string) bool {
	for _, r := range u.roles {
		if r.Name() == name {
			return true
		}
	}

	return false
}

func (u *User) Permissions() []Permission.Permission {
	var permissions []Permission.Permission
	for _, r := range u.roles {
		permissions = append(permissions, r.Permissions()...)
	}

	return permissions
}

func (u *User) HasPermission(name string) bool {
	for _, p := range u.Permissions() {
		if p.Name() == name {
			return true
		}
	}

	return false
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) ModifiedAt() time.Time {
	return u.modifiedAt
}

func (u *User) DeletedAt() time.Time {
	return u.deletedAt
}

func (u *User) IsDeleted() bool {
	return !u.DeletedAt().IsZero()
}

func (u *User) Delete() {
	u.deletedAt = time.Now()
	u.modifiedAt = u.deletedAt
}

func (u *User) Restore() {
	u.deletedAt = time.Time{}
	u.modifiedAt = time.Now()
}
