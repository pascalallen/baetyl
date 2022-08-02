package Permission

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Permission struct {
	gorm.Model
	Id          uuid.UUID `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
}

func Define(name string, description string) *Permission {
	id := uuid.New()
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
