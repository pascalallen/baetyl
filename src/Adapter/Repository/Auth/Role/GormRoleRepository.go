package Role

import (
	"errors"
	"fmt"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/Role"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type GormRoleRepository struct{}

func (repository GormRoleRepository) GetById(id ulid.ULID) (*Role.Role, error) {
	dsn := fmt.Sprintf(
		"dbname=%s user=%s password=%s host=%s port=%s",
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %s, error: %s", dsn, err.Error())
	}

	var role *Role.Role
	if err := db.Preload("Permissions").First(&role, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return role, nil
}

func (repository GormRoleRepository) GetByName(name string) (*Role.Role, error) {
	dsn := fmt.Sprintf(
		"dbname=%s user=%s password=%s host=%s port=%s",
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %s, error: %s", dsn, err.Error())
	}

	var role *Role.Role
	if err := db.Preload("Permissions").First(&role, "name = ?", name).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return role, nil
}

// GetAll TODO: Add pagination
func (repository GormRoleRepository) GetAll() (*[]Role.Role, error) {
	dsn := fmt.Sprintf(
		"dbname=%s user=%s password=%s host=%s port=%s",
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %s, error: %s", dsn, err.Error())
	}

	var roles *[]Role.Role

	if err := db.Find(&roles).Error; err != nil {
		return nil, fmt.Errorf("failed to get all roles, error: %s", err.Error())
	}

	return roles, nil
}

func (repository GormRoleRepository) Add(role *Role.Role) error {
	dsn := fmt.Sprintf(
		"dbname=%s user=%s password=%s host=%s port=%s",
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %s, error: %s", dsn, err.Error())
	}

	if err := db.Create(&role).Error; err != nil {
		return fmt.Errorf("failed to add role, error: %s", err.Error())
	}

	return nil
}

func (repository GormRoleRepository) Remove(role *Role.Role) error {
	dsn := fmt.Sprintf(
		"dbname=%s user=%s password=%s host=%s port=%s",
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %s, error: %s", dsn, err.Error())
	}

	if err := db.Delete(&role).Error; err != nil {
		return fmt.Errorf("failed to remove role, error: %s", err.Error())
	}

	return nil
}
