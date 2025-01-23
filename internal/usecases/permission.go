package usecases

import (
	"context"
	"gitlab.com/xamops/auth/internal/config"
	"gitlab.com/xamops/auth/internal/interfaces"
	"log/slog"
)

type PermissionCheckUsecases struct {
	config  *config.Config
	repo    interfaces.Repositories
	mappers interfaces.Mappers
	log     *slog.Logger
}

var _ interfaces.PermissionCheckUsecases = (*PermissionCheckUsecases)(nil)

func newPermissionCheckUsecases(deps Dependencies) *PermissionCheckUsecases {
	return &PermissionCheckUsecases{
		config:  deps.Config,
		repo:    deps.Repositories,
		mappers: deps.Mappers,
		log:     deps.Logger,
	}
}

func (p *PermissionCheckUsecases) OrganizationCheck(ctx context.Context, callerID, orgID string) (result bool, err error) {
	uid, err := p.mappers.UUID().FromString(callerID)
	if err != nil {
		return false, err
	}
	caller, err := p.repo.Users().ByID(p.repo.ConnWithContext(ctx), uid)
	if err != nil {
		return false, err
	}
	orgUUID, err := p.mappers.UUID().FromString(orgID)
	if err != nil {
		return false, err
	}
	if caller.OrganizationUUID == orgUUID {
		return true, nil
	}
	return
}

//func (a *AuthUsecases) CheckPermission(ctx context.Context, callerID string) error {
//	uid, err := a.mappers.UUID().FromString(callerID)
//	if err != nil {
//		return err
//	}
//	caller, err := a.repo.Users().ByID(a.repo.ConnWithContext(ctx), uid)
//	if err != nil {
//		return err
//	}
//	perm := a.mappers.Roles().PermissionsToClaims(caller.SystemRole.Permissions)
//
//	switch caller.SystemRole.Description {
//	case auth.DescriptionAdmin:
//		return fmt.Errorf("permission denied")
//	case auth.DescriptionOrganizationOwner:
//		if
//	case auth.DescriptionManager:
//	case auth.DescriptionUser:
//		if perm.Has()
//
//	}
//	panic("implement me")
//}
