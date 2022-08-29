package Permission

import (
	"errors"
	"fmt"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/Baetyl/src/Adapter/Database"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/Permission"
	"gorm.io/gorm"
)

type GormPermissionRepository struct{}

func (repository GormPermissionRepository) GetById(id ulid.ULID) (*Permission.Permission, error) {
	unitOfWork := Database.GormUnitOfWork{}
	if err := unitOfWork.InitDbSession(); err != nil {
		panic(err.Error())
	}

	db := unitOfWork.DatabaseSession

	var permission *Permission.Permission
	if err := db.First(permission, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return permission, nil
}

func (repository GormPermissionRepository) GetByName(name string) (*Permission.Permission, error) {
	unitOfWork := Database.GormUnitOfWork{}
	if err := unitOfWork.InitDbSession(); err != nil {
		panic(err.Error())
	}

	db := unitOfWork.DatabaseSession

	var permission *Permission.Permission
	if err := db.First(&permission, "name = ?", name).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return permission, nil
}

// GetAll TODO: Add pagination
func (repository GormPermissionRepository) GetAll() (*[]Permission.Permission, error) {
	unitOfWork := Database.GormUnitOfWork{}
	if err := unitOfWork.InitDbSession(); err != nil {
		panic(err.Error())
	}

	db := unitOfWork.DatabaseSession

	var permissions *[]Permission.Permission
	if err := db.Find(&permissions).Error; err != nil {
		return nil, fmt.Errorf("failed to get all permissions, error: %s", err.Error())
	}

	return permissions, nil
}

func (repository GormPermissionRepository) Add(permission *Permission.Permission) error {
	unitOfWork := Database.GormUnitOfWork{}
	if err := unitOfWork.InitDbSession(); err != nil {
		panic(err.Error())
	}

	db := unitOfWork.DatabaseSession

	if err := db.Create(&permission).Error; err != nil {
		return fmt.Errorf("failed to add permission, error: %s", err.Error())
	}

	return nil
}

func (repository GormPermissionRepository) Remove(permission *Permission.Permission) error {
	unitOfWork := Database.GormUnitOfWork{}
	if err := unitOfWork.InitDbSession(); err != nil {
		panic(err.Error())
	}

	db := unitOfWork.DatabaseSession

	if err := db.Delete(&permission).Error; err != nil {
		return fmt.Errorf("failed to remove permission, error: %s", err.Error())
	}

	return nil
}
