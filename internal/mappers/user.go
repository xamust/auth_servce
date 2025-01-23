package mappers

import (
	"gitlab.com/xamops/auth/internal/dto"
	"gitlab.com/xamops/auth/internal/entity"
	"gitlab.com/xamops/auth/internal/interfaces"
	"gitlab.com/xamops/auth/pkg/auth"
)

var _ interfaces.UsersMappers = (*UsersMappers)(nil)

type UsersMappers struct {
	role *RolesMappers
}

func newUsersMappers(role *RolesMappers) *UsersMappers {
	return &UsersMappers{
		role: role,
	}
}

func (u *UsersMappers) ToClaims(in *entity.User) *auth.UserClaims {
	return &auth.UserClaims{
		UUID:             in.UUID,
		OrganizationUUID: in.OrganizationUUID,
		Permissions:      *u.role.PermissionsToClaims(in.SystemRole.Permissions),
	}
}

func (u *UsersMappers) ToDTO(in *entity.User) (*dto.User, error) {
	role, err := u.role.ToDTO(in.SystemRole)
	if err != nil {
		return nil, err
	}
	return &dto.User{
		Email:          in.Email,
		OrganizationID: in.OrganizationUUID.String(),
		IsActive:       in.IsActive,
		Role:           role,
	}, nil
}
