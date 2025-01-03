package repositories

import (
	"gitlab.com/xamops/auth/internal/entity"
	"gitlab.com/xamops/auth/internal/interfaces"
	"gorm.io/gorm"
	"log/slog"
)

var _ interfaces.AuthRepository = (*AuthRepository)(nil)

type AuthRepository struct {
	log *slog.Logger
}

func newAuthRepository(log *slog.Logger) *AuthRepository {
	return &AuthRepository{
		log: log,
	}
}

func (a *AuthRepository) PasswordHashForUser(conn *gorm.DB, email string) (string, error) {
	var user entity.User
	if err := conn.First(&user, "email = ?", email).Error; err != nil {
		return "", err
	}
	return user.PasswordHash, nil
}

func (a *AuthRepository) ChangePassword(conn *gorm.DB, email string, hash string) error {
	if err := conn.
		Model(&entity.User{}).
		Where("email = ?", email).
		Update("password_hash", hash).
		Error; err != nil {
		return err
	}
	return nil
}
