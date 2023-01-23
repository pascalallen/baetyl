package User

import (
	"errors"
	"fmt"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/baetyl/src/Domain/Auth/User"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	UnitOfWork *gorm.DB
}

func (repository GormUserRepository) GetById(id ulid.ULID) (*User.User, error) {
	var user *User.User
	if err := repository.UnitOfWork.Preload("Roles.Permissions").First(&user, "id = ?", id.String()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to fetch User by ID: %s", id)
	}

	return user, nil
}

func (repository GormUserRepository) GetByEmailAddress(emailAddress string) (*User.User, error) {
	var user *User.User
	if err := repository.UnitOfWork.Preload("Roles.Permissions").First(&user, "email_address = ?", emailAddress).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to fetch User by email address: %s", emailAddress)
	}

	return user, nil
}

// GetAll TODO: Add pagination
func (repository GormUserRepository) GetAll(includeDeleted bool) (*[]User.User, error) {
	var users *[]User.User
	if !includeDeleted {
		repository.UnitOfWork = repository.UnitOfWork.Where("deleted_at IS NULL")
	}

	if err := repository.UnitOfWork.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch all Users: %s", err)
	}

	return users, nil
}

func (repository GormUserRepository) Add(user *User.User) error {
	if err := repository.UnitOfWork.Create(&user).Error; err != nil {
		return fmt.Errorf("failed to persist User to database: %s", user)
	}

	return nil
}

func (repository GormUserRepository) Remove(user *User.User) error {
	user.Delete()

	if err := repository.UnitOfWork.Save(&user).Error; err != nil {
		return fmt.Errorf("failed to delete User from database: %s", user)
	}

	return nil
}

func (repository GormUserRepository) UpdateOrAdd(user *User.User) error {
	if err := repository.UnitOfWork.Save(&user).Error; err != nil {
		return fmt.Errorf("failed to update User: %s", user)
	}

	return nil
}
