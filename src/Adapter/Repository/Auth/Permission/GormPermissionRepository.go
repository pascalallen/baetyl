package Permission

import (
	"errors"
	"fmt"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/Permission"
	"gorm.io/gorm"
)

type GormPermissionRepository struct {
	DatabaseConnection *gorm.DB
}

func (repository GormPermissionRepository) GetById(id ulid.ULID) (*Permission.Permission, error) {
	var permission *Permission.Permission
	if err := repository.DatabaseConnection.First(&permission, "id = ?", id.String()).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return permission, nil
}

func (repository GormPermissionRepository) GetByName(name string) (*Permission.Permission, error) {
	var permission *Permission.Permission
	if err := repository.DatabaseConnection.First(&permission, "name = ?", name).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return permission, nil
}

// GetAll TODO: Add pagination
func (repository GormPermissionRepository) GetAll() (*[]Permission.Permission, error) {
	var permissions *[]Permission.Permission
	if err := repository.DatabaseConnection.Find(&permissions).Error; err != nil {
		return nil, fmt.Errorf("failed to get all permissions, error: %s", err.Error())
	}

	return permissions, nil
}

func (repository GormPermissionRepository) Add(permission *Permission.Permission) error {
	if err := repository.DatabaseConnection.Create(&permission).Error; err != nil {
		return fmt.Errorf("failed to add permission, error: %s", err.Error())
	}

	return nil
}

func (repository GormPermissionRepository) Remove(permission *Permission.Permission) error {
	if err := repository.DatabaseConnection.Delete(&permission).Error; err != nil {
		return fmt.Errorf("failed to remove permission, error: %s", err.Error())
	}

	return nil
}

func (repository GormPermissionRepository) Save(permission *Permission.Permission) error {
	if err := repository.DatabaseConnection.Save(&permission).Error; err != nil {
		return fmt.Errorf("failed to save permission, error: %s", err.Error())
	}

	return nil
}
