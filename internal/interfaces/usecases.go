package interfaces

import (
	"context"
	"github.com/google/uuid"
	"gitlab.com/xamops/auth/internal/entity"
)

type Usecases interface {
	Auth() AuthUsecases
	Users() UsersUsecases
	Roles() SystemRolesUsecases
	Permissions() PermissionCheckUsecases
}

type AuthUsecases interface {
	ChangePassword(ctx context.Context, email string, passwordOld string, passwordNew string) error
	Login(ctx context.Context, email string, password string) (*entity.User, error)
}

type UsersUsecases interface {
	GetByID(ctx context.Context, id string) (*entity.User, error)
	GetByUUID(ctx context.Context, uid uuid.UUID) (*entity.User, error)
}

type SystemRolesUsecases interface {
	GetByID(ctx context.Context, id string) (*entity.SystemRole, error)
	GetByUUID(ctx context.Context, uid uuid.UUID) (*entity.SystemRole, error)
	List(ctx context.Context, limit, offset int) ([]entity.SystemRole, error)
}

type PermissionCheckUsecases interface {
	OrganizationCheck(ctx context.Context, callerID, orgID string) (bool, error)
}
