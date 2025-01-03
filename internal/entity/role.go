package entity

import (
	"github.com/google/uuid"
	"gitlab.com/xamops/auth/pkg/metadata"
)

type Role struct {
	Common
	Permission     Permission
	Description    string
	OrganizationID uuid.UUID
	Metadata       metadata.Metadata
}
