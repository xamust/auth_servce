package mappers

import (
	"gitlab.com/xamops/auth/internal/dto"
	"gitlab.com/xamops/auth/internal/entity"
	"gitlab.com/xamops/auth/internal/interfaces"
	"gitlab.com/xamops/auth/pkg/auth"
)

const (
	role_title       = "role_title"
	permission_title = "permission_title"
)

var _ interfaces.RolesMappers = (*RolesMappers)(nil)

type RolesMappers struct {
	metadataMappers *MetadataMappers
}

func newRolesMappers(metadataMappers *MetadataMappers) *RolesMappers {
	return &RolesMappers{
		metadataMappers: metadataMappers,
	}
}

func (r *RolesMappers) ToDTO(in *entity.SystemRole) (*dto.Role, error) {
	metadataTitle, err := r.metadataMappers.ByKey(role_title, in.Metadata)
	if err != nil {
		return nil, err
	}
	permissions, err := r.PermissionsToDTO(in.Permissions)
	if err != nil {
		return nil, err
	}
	return &dto.Role{
		Description: string(in.Description),
		Title:       metadataTitle.(string),
		Permissions: permissions,
	}, nil
}

func (r *RolesMappers) PermissionsToClaims(in []entity.SystemRolesPermissionsRelation) *auth.Permissions {
	perm := make([]auth.Permission, len(in))
	for i, relation := range in {
		perm[i] = auth.Permission(relation.SystemRolesPermissions.Description)
	}
	return auth.NewPermissions(perm...)
}

func (r *RolesMappers) PermissionsToDTO(in []entity.SystemRolesPermissionsRelation) ([]dto.Permission, error) {
	result := make([]dto.Permission, len(in))
	for i, relation := range in {
		metadataTitle, err := r.metadataMappers.ByKey(permission_title, relation.SystemRolesPermissions.Metadata)
		if err != nil {
			return nil, err
		}
		result[i] = dto.Permission{
			Description: string(relation.SystemRolesPermissions.Description),
			Title:       metadataTitle.(string),
		}
	}
	return result, nil
}
