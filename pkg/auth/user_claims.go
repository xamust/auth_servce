package auth

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type UserClaims struct {
	UUID             uuid.UUID   `json:"uuid"`
	OrganizationUUID uuid.UUID   `json:"organization_uuid"`
	EmployeeUUID     uuid.UUID   `json:"employee_uuid"`
	Permissions      Permissions `json:"permissions"`
	jwt.StandardClaims
}

var _ jwt.Claims = (*UserClaims)(nil)
