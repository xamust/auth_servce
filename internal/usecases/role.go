package usecases

import (
	"context"
	"github.com/google/uuid"
	"gitlab.com/xamops/auth/internal/config"
	"gitlab.com/xamops/auth/internal/entity"
	"gitlab.com/xamops/auth/internal/interfaces"
	"log/slog"
)

type SystemRolesUsecases struct {
	config  *config.Config
	repo    interfaces.Repositories
	mappers interfaces.Mappers
	log     *slog.Logger
}

var _ interfaces.SystemRolesUsecases = (*SystemRolesUsecases)(nil)

func newSystemRolesUsecases(deps Dependencies) *SystemRolesUsecases {
	return &SystemRolesUsecases{
		config:  deps.Config,
		repo:    deps.Repositories,
		mappers: deps.Mappers,
		log:     deps.Logger,
	}
}

func (a *SystemRolesUsecases) GetByID(ctx context.Context, id string) (*entity.SystemRole, error) {
	panic("implement me")
}

func (a *SystemRolesUsecases) GetByUUID(ctx context.Context, uid uuid.UUID) (*entity.SystemRole, error) {
	panic("implement me")
}

func (a *SystemRolesUsecases) List(ctx context.Context, limit, offset int) ([]entity.SystemRole, error) {
	//a.repo.Roles().ListByOrganizationID()
	panic("implement me")
}
