package entity

import (
	"github.com/google/uuid"
	"gitlab.com/xamops/auth/pkg/metadata"
)

type User struct {
	Common
	Email            string
	PasswordHash     string
	OrganizationUUID uuid.UUID
	IsActive         bool
	Metadata         metadata.Metadata
	SystemRoleUUID   uuid.UUID
	SystemRole       *SystemRole `gorm:"foreignKey:SystemRoleUUID;references:UUID"`
}
