package Role

import (
	"errors"
	"fmt"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/Role"
	"gorm.io/gorm"
)

type GormRoleRepository struct {
	DatabaseConnection *gorm.DB
}

func (repository GormRoleRepository) GetById(id ulid.ULID) (*Role.Role, error) {
	var role *Role.Role
	if err := repository.DatabaseConnection.Preload("Permissions").First(&role, "id = ?", id.String()).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return role, nil
}

func (repository GormRoleRepository) GetByName(name string) (*Role.Role, error) {
	var role *Role.Role
	if err := repository.DatabaseConnection.Preload("Permissions").First(&role, "name = ?", name).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return role, nil
}

// GetAll TODO: Add pagination
func (repository GormRoleRepository) GetAll() (*[]Role.Role, error) {
	var roles *[]Role.Role
	if err := repository.DatabaseConnection.Find(&roles).Error; err != nil {
		return nil, fmt.Errorf("failed to get all roles, error: %s", err.Error())
	}

	return roles, nil
}

func (repository GormRoleRepository) Add(role *Role.Role) error {
	if err := repository.DatabaseConnection.Create(&role).Error; err != nil {
		return fmt.Errorf("failed to add role, error: %s", err.Error())
	}

	return nil
}

func (repository GormRoleRepository) Remove(role *Role.Role) error {
	if err := repository.DatabaseConnection.Delete(&role).Error; err != nil {
		return fmt.Errorf("failed to remove role, error: %s", err.Error())
	}

	return nil
}

func (repository GormRoleRepository) Save(role *Role.Role) error {
	if err := repository.DatabaseConnection.Save(&role).Error; err != nil {
		return fmt.Errorf("failed to save role, error: %s", err.Error())
	}

	return nil
}
