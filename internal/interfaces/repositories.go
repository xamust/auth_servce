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
	Roles() SystemRolesRepository
	ConnWithContext(ctx context.Context) *gorm.DB
	Conn() *gorm.DB
}

type AuthRepository interface {
	PasswordHashForUser(conn *gorm.DB, email string) (string, error)
	ChangePassword(conn *gorm.DB, email string, hash string) error
}

type UsersRepository interface {
	ByID(db *gorm.DB, uid uuid.UUID) (*entity.User, error)
	ByEmail(db *gorm.DB, email string) (*entity.User, error)
	ListByOrganizationID(db *gorm.DB, orgID uuid.UUID, limit, offset int) ([]entity.User, error)
	Create(db *gorm.DB, user *entity.User) (*entity.User, error)
	Update(db *gorm.DB, user *entity.User) (*entity.User, error)
	DeleteByID(db *gorm.DB, uid uuid.UUID) error
}

type SystemRolesRepository interface {
	ByID(db *gorm.DB, uid uuid.UUID) (*entity.SystemRole, error)
	ByIDWithOrgnizationID(db *gorm.DB, uid uuid.UUID, orgID uuid.UUID) (*entity.SystemRole, error)
	ListByOrganizationID(db *gorm.DB, orgID uuid.UUID, limit, offset int) ([]entity.SystemRole, error)
}
