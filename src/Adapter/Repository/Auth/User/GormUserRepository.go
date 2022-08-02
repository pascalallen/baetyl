package User

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/User"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type GormUserRepository struct{}

func (repository GormUserRepository) GetById(id uuid.UUID) *User.User {
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
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
		panic("failed to connect database")
	}

	var user *User.User
	db.Preload("Roles.Permissions").First(&user, id)

	return user
}

func (repository GormUserRepository) GetByEmailAddress(emailAddress string) *User.User {
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
		panic("failed to connect database")
	}

	var user *User.User
	db.Preload("Roles.Permissions").First(&user, "email_address = ?", emailAddress)

	return user
}

func (repository GormUserRepository) GetAll(includeDeleted bool) *[]User.User {
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
		panic("failed to connect database")
	}

	var users *[]User.User
	if !includeDeleted {
		db = db.Where("deleted_at IS NULL")
	}

	db.Find(&users)

	return users
}

func (repository GormUserRepository) Add(user *User.User) {
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
		panic("failed to connect database")
	}

	db.Create(&user)
}

func (repository GormUserRepository) Remove(user *User.User) {
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
		panic("failed to connect database")
	}

	user.Delete()

	db.Save(&user)
}
