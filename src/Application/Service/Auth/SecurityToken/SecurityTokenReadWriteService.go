package SecurityToken

import (
	"errors"
	"fmt"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/Crypto"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/SecurityToken"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/User"
	"gorm.io/gorm"
	"time"
)

type SecurityTokenReadWriteService struct {
	SecurityTokenRepository SecurityToken.SecurityTokenRepository
	UnitOfWork              *gorm.DB
}

func (service SecurityTokenReadWriteService) GetById(id ulid.ULID) (*SecurityToken.SecurityToken, error) {
	var securityToken *SecurityToken.SecurityToken
	if err := service.UnitOfWork.Preload("User.Roles.Permissions").First(&securityToken, "id = ?", id.String()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to fetch SecurityToken by ID: %s", id)
	}

	return securityToken, nil
}

func (service SecurityTokenReadWriteService) GetByCrypto(crypto Crypto.Crypto) (*SecurityToken.SecurityToken, error) {
	var securityToken *SecurityToken.SecurityToken
	if err := service.UnitOfWork.Preload("User.Roles.Permissions").First(&securityToken, "crypto = ?", crypto).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to fetch SecurityToken by crypto: %s", crypto)
	}

	return securityToken, nil
}

func (service SecurityTokenReadWriteService) GetAllForUser(user User.User) (*[]SecurityToken.SecurityToken, error) {
	var securityTokens *[]SecurityToken.SecurityToken
	if err := service.UnitOfWork.Preload("User.Roles.Permissions").Where("user_id = ?", user.Id).Find(&securityTokens).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch all SecurityTokens for user: %s", err)
	}

	return securityTokens, nil
}

func (service SecurityTokenReadWriteService) Add(securityToken *SecurityToken.SecurityToken) error {
	if err := service.UnitOfWork.Create(&securityToken).Error; err != nil {
		return fmt.Errorf("failed to persist SecurityToken to database: %s", securityToken)
	}

	return nil
}

func (service SecurityTokenReadWriteService) Remove(securityToken *SecurityToken.SecurityToken) error {
	if err := service.UnitOfWork.Delete(&securityToken).Error; err != nil {
		return fmt.Errorf("failed to delete SecurityToken from database: %s", securityToken)
	}

	return nil
}

func (service SecurityTokenReadWriteService) UpdateOrAdd(securityToken *SecurityToken.SecurityToken) error {
	if err := service.UnitOfWork.Save(&securityToken).Error; err != nil {
		return fmt.Errorf("failed to update SecurityToken: %s", securityToken)
	}

	return nil
}

func (service SecurityTokenReadWriteService) ClearExpiredTokens() error {
	if err := service.UnitOfWork.Delete(&SecurityToken.SecurityToken{}, "expires_at <= ?", time.Now()).Error; err != nil {
		return fmt.Errorf("failed to clear expired SecurityTokens from database: %s", err)
	}

	return nil
}
