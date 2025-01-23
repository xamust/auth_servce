package interfaces

import (
	"github.com/google/uuid"
	"gitlab.com/xamops/auth/internal/dto"
	"gitlab.com/xamops/auth/internal/entity"
	"gitlab.com/xamops/auth/pkg/auth"
)

type Mappers interface {
	UUID() UUIDMappers
	Users() UsersMappers
	Roles() RolesMappers
	Metadata() MetadataMappers
}

type UUIDMappers interface {
	FromString(uid string) (uuid.UUID, error)
	ToString(id uuid.UUID) string
}

type UsersMappers interface {
	ToClaims(in *entity.User) *auth.UserClaims
	ToDTO(in *entity.User) (*dto.User, error)
}

type RolesMappers interface {
	ToDTO(in *entity.SystemRole) (*dto.Role, error)
	PermissionsToClaims(in []entity.SystemRolesPermissionsRelation) *auth.Permissions
	PermissionsToDTO(in []entity.SystemRolesPermissionsRelation) ([]dto.Permission, error)
}

type MetadataMappers interface {
	ByKey(key string, in []byte) (interface{}, error)
	FromData(in []byte) (map[string]interface{}, error)
	ToData(in map[string]interface{}) ([]byte, error)
}
