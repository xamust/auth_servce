package entity

import (
	"github.com/google/uuid"
	"gitlab.com/xamops/auth/pkg/metadata"
)

type Permission struct {
	Common
	Description    string
	OrganizationID uuid.UUID
	Metadata       metadata.Metadata
}
