package repositories

import (
	"github.com/google/uuid"
	"gitlab.com/xamops/auth/internal/entity"
	"gitlab.com/xamops/auth/internal/interfaces"
	"gorm.io/gorm"
	"log/slog"
)

var _ interfaces.SystemRolesRepository = (*SystemRolesRepository)(nil)

func newSystemRolesRepository(log *slog.Logger) *SystemRolesRepository {
	return &SystemRolesRepository{
		log: log,
	}
}

type SystemRolesRepository struct {
	log *slog.Logger
}

func (s *SystemRolesRepository) ByID(db *gorm.DB, uid uuid.UUID) (*entity.SystemRole, error) {
	var role entity.SystemRole
	err := db.
		Preload("Permissions.SystemRolesPermissions").
		Where("uuid = ?", uid.String()).
		First(&role).
		Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (s *SystemRolesRepository) ByIDWithOrgnizationID(db *gorm.DB, uid uuid.UUID, orgID uuid.UUID) (*entity.SystemRole, error) {
	var role entity.SystemRole
	err := db.
		Preload("Permissions.SystemRolesPermissions").
		Where("uuid = ?", uid.String()).
		Where("organization_uuid = ?", orgID.String()).
		First(&role).
		Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (s *SystemRolesRepository) ListByOrganizationID(db *gorm.DB, orgID uuid.UUID, limit, offset int) ([]entity.SystemRole, error) {
	var roles []entity.SystemRole
	err := db.
		Preload("Permissions.SystemRolesPermissions").
		Where("organization_uuid = ?", orgID.String()).
		Limit(limit).
		Offset(offset).
		Find(roles).
		Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}
