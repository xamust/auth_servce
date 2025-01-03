package entity

import (
	"github.com/google/uuid"
	"gitlab.com/xamops/auth/pkg/metadata"
)

type User struct {
	Common
	Email          string
	PasswordHash   string
	OrganizationID uuid.UUID
	IsActive       bool
	Metadata       metadata.Metadata
	Role           *Role
}
