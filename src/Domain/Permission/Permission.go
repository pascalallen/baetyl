package Permission

import (
	"github.com/google/uuid"
	"time"
)

type Permission struct {
	id          uuid.UUID `validate:"required,uuid"`
	name        string    `validate:"required"`
	description string    `validate:"required"`
	createdAt   time.Time `validate:"required,datetime"`
	modifiedAt  time.Time `validate:"required,datetime"`
}

func Define(name string, description string) Permission {
	id := uuid.New()
	createdAt := time.Now()

	return Permission{
		id:          id,
		name:        name,
		description: description,
		createdAt:   createdAt,
		modifiedAt:  createdAt,
	}
}

func (p Permission) Id() uuid.UUID {
	return p.id
}

func (p Permission) Name() string {
	return p.name
}

func (p Permission) UpdateName(name string) {
	p.name = name
	p.modifiedAt = time.Now()
}

func (p Permission) Description() string {
	return p.description
}

func (p Permission) UpdateDescription(description string) {
	p.description = description
	p.modifiedAt = time.Now()
}

func (p Permission) CreatedAt() time.Time {
	return p.createdAt
}

func (p Permission) ModifiedAt() time.Time {
	return p.modifiedAt
}
