package Permission

import (
	"errors"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/Permission"
	"gorm.io/gorm"
)

type GormPermissionRepository struct {
	UnitOfWork *gorm.DB
}

func (repository GormPermissionRepository) GetById(id ulid.ULID) (*Permission.Permission, error) {
	var permission *Permission.Permission
	if err := repository.UnitOfWork.First(&permission, "id = ?", id.String()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return permission, nil
}

func (repository GormPermissionRepository) GetByName(name string) (*Permission.Permission, error) {
	var permission *Permission.Permission
	if err := repository.UnitOfWork.First(&permission, "name = ?", name).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return permission, nil
}

// GetAll TODO: Add pagination
func (repository GormPermissionRepository) GetAll() (*[]Permission.Permission, error) {
	var permissions *[]Permission.Permission
	if err := repository.UnitOfWork.Find(&permissions).Error; err != nil {
		return nil, err
	}

	return permissions, nil
}

func (repository GormPermissionRepository) Add(permission *Permission.Permission) error {
	if err := repository.UnitOfWork.Create(&permission).Error; err != nil {
		return err
	}

	return nil
}

func (repository GormPermissionRepository) Remove(permission *Permission.Permission) error {
	if err := repository.UnitOfWork.Delete(&permission).Error; err != nil {
		return err
	}

	return nil
}

func (repository GormPermissionRepository) Save(permission *Permission.Permission) error {
	if err := repository.UnitOfWork.Save(&permission).Error; err != nil {
		return err
	}

	return nil
}
