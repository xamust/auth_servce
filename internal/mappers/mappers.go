package mappers

import (
	"gitlab.com/xamops/auth/internal/interfaces"
)

type Mappers struct {
	uuidMapper     *UUIDMappers
	userMapper     *UsersMappers
	roleMapper     *RolesMappers
	metadataMapper *MetadataMappers
}

func New() interfaces.Mappers {
	metadata := newMetadataMappers()
	role := newRolesMappers(metadata)
	return &Mappers{
		uuidMapper:     newUUIDMappers(),
		userMapper:     newUsersMappers(role),
		roleMapper:     role,
		metadataMapper: metadata,
	}
}

func (m *Mappers) Users() interfaces.UsersMappers {
	return m.userMapper
}

func (m *Mappers) UUID() interfaces.UUIDMappers {
	return m.uuidMapper
}

func (m *Mappers) Roles() interfaces.RolesMappers {
	return m.roleMapper
}

func (m *Mappers) Metadata() interfaces.MetadataMappers {
	return m.metadataMapper
}
