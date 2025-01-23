package entity

import (
	"github.com/google/uuid"
	"gitlab.com/xamops/auth/pkg/auth"
	"gitlab.com/xamops/auth/pkg/metadata"
	"gorm.io/gorm"
	"time"
)

type SystemRole struct {
	Common
	Description      auth.DescriptionSystemRole
	OrganizationUUID uuid.UUID
	Metadata         metadata.Metadata
	Permissions      []SystemRolesPermissionsRelation `gorm:"foreignKey:SystemRoleUUID;references:UUID"`
}

type SystemRolesPermissions struct {
	Common
	Description    auth.Permission
	OrganizationID uuid.UUID
	Metadata       metadata.Metadata
}

type SystemRolesPermissionsRelation struct {
	UUID                   string `gorm:"type:uuid;primaryKey"`
	SystemRoleUUID         string `gorm:"type:uuid;not null"`
	PermissionUUID         string `gorm:"type:uuid;not null"`
	OrganizationUUID       string `gorm:"type:uuid;not null"`
	CreatedAt              time.Time
	UpdatedAt              time.Time
	DeletedAt              gorm.DeletedAt          `gorm:"index"`
	SystemRolesPermissions *SystemRolesPermissions `gorm:"foreignKey:PermissionUUID;references:UUID"`
}
