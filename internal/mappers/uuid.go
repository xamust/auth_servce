package mappers

import (
	"github.com/google/uuid"
	"gitlab.com/xamops/auth/internal/interfaces"
)

var _ interfaces.UUIDMappers = (*UUIDMappers)(nil)

type UUIDMappers struct {
}

func newUUIDMappers() *UUIDMappers {
	return &UUIDMappers{}
}

func (U *UUIDMappers) FromString(uid string) (uuid.UUID, error) {
	return uuid.Parse(uid)
}

func (U *UUIDMappers) ToString(id uuid.UUID) string {
	return id.String()
}
