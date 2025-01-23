package dto

type User struct {
	Email          string `json:"email"`
	OrganizationID string `json:"organization_id"`
	IsActive       bool   `json:"is_active"`
	Role           *Role  `json:"role"`
}
