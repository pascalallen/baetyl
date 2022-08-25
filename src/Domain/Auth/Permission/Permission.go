package Permission

import (
	"github.com/oklog/ulid/v2"
	"time"
)

type Permission struct {
	Id          ulid.ULID `json:"id" gorm:"primaryKey;size:26;not null"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
}

func Define(id ulid.ULID, name string, description string) *Permission {
	createdAt := time.Now()

	return &Permission{
		Id:          id,
		Name:        name,
		Description: description,
		CreatedAt:   createdAt,
		ModifiedAt:  createdAt,
	}
}

func (p *Permission) UpdateName(name string) {
	p.Name = name
	p.ModifiedAt = time.Now()
}

func (p *Permission) UpdateDescription(description string) {
	p.Description = description
	p.ModifiedAt = time.Now()
}
