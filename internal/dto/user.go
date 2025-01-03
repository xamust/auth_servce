package dto

import "gitlab.com/xamops/auth/internal/entity"

type User struct {
	Email          string      `json:"email"`
	OrganizationID string      `json:"organization_id"`
	IsActive       bool        `json:"is_active"`
	Role           entity.Role `json:"role"`
}
