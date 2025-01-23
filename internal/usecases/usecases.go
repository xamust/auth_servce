package usecases

import (
	"gitlab.com/xamops/auth/internal/config"
	"gitlab.com/xamops/auth/internal/interfaces"
	"log/slog"
)

var _ interfaces.Usecases = (*Usecases)(nil)

type Dependencies struct {
	Config       *config.Config
	Repositories interfaces.Repositories
	Mappers      interfaces.Mappers
	Logger       *slog.Logger
}

func New(deps Dependencies) *Usecases {
	return &Usecases{
		auth:       newAuthUsecases(deps),
		user:       newUsersUsecases(deps),
		role:       newSystemRolesUsecases(deps),
		permission: newPermissionCheckUsecases(deps),
	}
}

type Usecases struct {
	auth       interfaces.AuthUsecases
	user       interfaces.UsersUsecases
	role       interfaces.SystemRolesUsecases
	permission interfaces.PermissionCheckUsecases
}

func (u *Usecases) Auth() interfaces.AuthUsecases {
	return u.auth
}

func (u *Usecases) Users() interfaces.UsersUsecases {
	return u.user
}

func (u *Usecases) Roles() interfaces.SystemRolesUsecases {
	return u.role
}

func (u *Usecases) Permissions() interfaces.PermissionCheckUsecases {
	return u.permission
}
