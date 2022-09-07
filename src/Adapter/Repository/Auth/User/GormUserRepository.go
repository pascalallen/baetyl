package User

import (
	"errors"
	"fmt"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/User"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	DatabaseConnection *gorm.DB
}

func (repository GormUserRepository) GetById(id ulid.ULID) (*User.User, error) {
	var user *User.User
	if err := repository.DatabaseConnection.Preload("Roles.Permissions").First(&user, "id = ?", id.String()).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return user, nil
}

func (repository GormUserRepository) GetByEmailAddress(emailAddress string) (*User.User, error) {
	var user *User.User
	if err := repository.DatabaseConnection.Preload("Roles.Permissions").First(&user, "email_address = ?", emailAddress).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return user, nil
}

// GetAll TODO: Add pagination
func (repository GormUserRepository) GetAll(includeDeleted bool) (*[]User.User, error) {
	var users *[]User.User
	if !includeDeleted {
		repository.DatabaseConnection = repository.DatabaseConnection.Where("deleted_at IS NULL")
	}

	if err := repository.DatabaseConnection.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("failed to get all users, error: %s", err.Error())
	}

	return users, nil
}

func (repository GormUserRepository) Add(user *User.User) error {
	if err := repository.DatabaseConnection.Create(&user).Error; err != nil {
		return fmt.Errorf("failed to add user, error: %s", err.Error())
	}

	return nil
}

func (repository GormUserRepository) Remove(user *User.User) error {
	user.Delete()

	if err := repository.DatabaseConnection.Save(&user).Error; err != nil {
		return fmt.Errorf("failed to remove user, error: %s", err.Error())
	}

	return nil
}

func (repository GormUserRepository) Save(user *User.User) error {
	if err := repository.DatabaseConnection.Save(&user).Error; err != nil {
		return fmt.Errorf("failed to save user, error: %s", err.Error())
	}

	return nil
}
