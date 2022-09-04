package User

import (
	"errors"
	"fmt"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/User"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type GormUserRepository struct{}

func (repository GormUserRepository) GetById(id ulid.ULID) (*User.User, error) {
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

	var user *User.User
	if err := db.Preload("Roles.Permissions").First(&user, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return user, nil
}

func (repository GormUserRepository) GetByEmailAddress(emailAddress string) (*User.User, error) {
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

	var user *User.User
	if err := db.Preload("Roles.Permissions").First(&user, "email_address = ?", emailAddress).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return user, nil
}

// GetAll TODO: Add pagination
func (repository GormUserRepository) GetAll(includeDeleted bool) (*[]User.User, error) {
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

	var users *[]User.User
	if !includeDeleted {
		db = db.Where("deleted_at IS NULL")
	}

	if err := db.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("failed to get all users, error: %s", err.Error())
	}

	return users, nil
}

func (repository GormUserRepository) Add(user *User.User) error {
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

	if err := db.Create(&user).Error; err != nil {
		return fmt.Errorf("failed to add user, error: %s", err.Error())
	}

	return nil
}

func (repository GormUserRepository) Remove(user *User.User) error {
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

	user.Delete()

	if err := db.Save(&user).Error; err != nil {
		return fmt.Errorf("failed to remove user, error: %s", err.Error())
	}

	return nil
}
