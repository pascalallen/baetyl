package Role

import (
	"errors"
	"fmt"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/baetyl/src/Domain/Auth/Role"
	"gorm.io/gorm"
)

type GormRoleRepository struct {
	UnitOfWork *gorm.DB
}

func (repository GormRoleRepository) GetById(id ulid.ULID) (*Role.Role, error) {
	var role *Role.Role
	if err := repository.UnitOfWork.Preload("Permissions").First(&role, "id = ?", id.String()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to fetch Role by ID: %s", id)
	}

	return role, nil
}

func (repository GormRoleRepository) GetByName(name string) (*Role.Role, error) {
	var role *Role.Role
	if err := repository.UnitOfWork.Preload("Permissions").First(&role, "name = ?", name).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to fetch Role by name: %s", name)
	}

	return role, nil
}

// GetAll TODO: Add pagination
func (repository GormRoleRepository) GetAll() (*[]Role.Role, error) {
	var roles *[]Role.Role
	if err := repository.UnitOfWork.Find(&roles).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch all Roles: %s", err)
	}

	return roles, nil
}

func (repository GormRoleRepository) Add(role *Role.Role) error {
	if err := repository.UnitOfWork.Create(&role).Error; err != nil {
		return fmt.Errorf("failed to persist Role to database: %s", role)
	}

	return nil
}

func (repository GormRoleRepository) Remove(role *Role.Role) error {
	if err := repository.UnitOfWork.Delete(&role).Error; err != nil {
		return fmt.Errorf("failed to delete Role from database: %s", role)
	}

	return nil
}

func (repository GormRoleRepository) UpdateOrAdd(role *Role.Role) error {
	if err := repository.UnitOfWork.Save(&role).Error; err != nil {
		return fmt.Errorf("failed to update Role: %s", role)
	}

	return nil
}
