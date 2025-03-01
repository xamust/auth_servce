package repositories

import (
	"github.com/google/uuid"
	"gitlab.com/xamops/auth/internal/entity"
	"gitlab.com/xamops/auth/internal/interfaces"
	"gorm.io/gorm"
	"log/slog"
)

var _ interfaces.UsersRepository = (*UsersRepository)(nil)

func newUsersRepository(log *slog.Logger) *UsersRepository {
	return &UsersRepository{
		log: log,
	}
}

type UsersRepository struct {
	log *slog.Logger
}

func (u *UsersRepository) ByID(db *gorm.DB, uid uuid.UUID) (*entity.User, error) {
	var user entity.User
	err := db.
		Preload("SystemRole.Permissions.SystemRolesPermissions").
		Where("uuid = ?", uid.String()).
		First(&user).
		Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UsersRepository) ByEmail(db *gorm.DB, email string) (*entity.User, error) {
	var user entity.User
	err := db.
		Preload("SystemRole.Permissions.SystemRolesPermissions").
		Where("email = ?", email).
		First(&user).
		Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UsersRepository) ListByOrganizationID(db *gorm.DB, orgID uuid.UUID, limit, offset int) ([]entity.User, error) {
	panic("implement me")
}

func (u *UsersRepository) Create(db *gorm.DB, user *entity.User) (*entity.User, error) {
	panic("implement me")
}

func (u *UsersRepository) Update(db *gorm.DB, user *entity.User) (*entity.User, error) {
	panic("implement me")
}

func (u *UsersRepository) DeleteByID(db *gorm.DB, uid uuid.UUID) error {
	panic("implement me")
}
