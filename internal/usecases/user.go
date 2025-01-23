package usecases

import (
	"context"
	"github.com/google/uuid"
	"gitlab.com/xamops/auth/internal/config"
	"gitlab.com/xamops/auth/internal/entity"
	"gitlab.com/xamops/auth/internal/interfaces"
	"log/slog"
)

type UsersUsecases struct {
	config  *config.Config
	repo    interfaces.Repositories
	mappers interfaces.Mappers
	log     *slog.Logger
}

var _ interfaces.UsersUsecases = (*UsersUsecases)(nil)

func newUsersUsecases(deps Dependencies) *UsersUsecases {
	return &UsersUsecases{
		config:  deps.Config,
		repo:    deps.Repositories,
		mappers: deps.Mappers,
		log:     deps.Logger,
	}
}

func (u *UsersUsecases) GetByID(ctx context.Context, id string) (*entity.User, error) {
	usrUUID, err := u.mappers.UUID().FromString(id)
	if err != nil {
		return nil, err
	}
	return u.repo.Users().ByID(u.repo.ConnWithContext(ctx), usrUUID)
}

func (u *UsersUsecases) GetByUUID(ctx context.Context, uid uuid.UUID) (*entity.User, error) {
	return u.repo.Users().ByID(u.repo.ConnWithContext(ctx), uid)
}
