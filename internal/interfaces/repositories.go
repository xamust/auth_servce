package interfaces

import (
	"context"
	"github.com/google/uuid"
	"gitlab.com/xamops/auth/internal/entity"
	"gorm.io/gorm"
)

type Repositories interface {
	Auth() AuthRepository
	Users() UsersRepository
	ConnWithContext(ctx context.Context) *gorm.DB
	Conn() *gorm.DB
}

type AuthRepository interface {
	PasswordHashForUser(conn *gorm.DB, email string) (string, error)
	ChangePassword(conn *gorm.DB, email string, hash string) error
}

type UsersRepository interface {
	GetByID(db *gorm.DB, uid uuid.UUID) (*entity.User, error)
	ListByOrganizationID(db *gorm.DB, orgID uuid.UUID) ([]entity.User, error)
	Create(db *gorm.DB, user *entity.User) (*entity.User, error)
	Update(db *gorm.DB, user *entity.User) (*entity.User, error)
	DeleteByID(db *gorm.DB, uid uuid.UUID) error
}
