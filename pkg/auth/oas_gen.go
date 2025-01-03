// Code generated by tools/gen_permissions. DO NOT EDIT.
package auth

// Код разрешения.
// Ref: #/components/schemas/Permission
type Permission string

const (
	PermissionManageLeaders Permission = "manage_leaders"
	PermissionManageEmployees Permission = "manage_employees"
	PermissionManageUnits Permission = "manage_units"
	PermissionManageMatrices Permission = "manage_matrices"
	PermissionConductReview Permission = "conduct_review"
	PermissionViewEmployees Permission = "view_employees"
	PermissionChangePassword Permission = "change_password"
	PermissionManageOrganizations Permission = "manage_organizations"
)

// Код роли.
// Ref: #/components/schemas/Codename
type Codename string

const (
	CodenameAdmin Codename = "admin"
	CodenameUnitLeader Codename = "unit_leader"
	CodenameReviewer Codename = "reviewer"
	CodenameUser Codename = "user"
)

