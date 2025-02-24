package auth

type Permission string

const (
	PermissionManageAllUsers                Permission = "manage_all_users"
	PermissionManageAllRoles                Permission = "manage_all_roles"
	PermissionManageAllPermissions          Permission = "manage_all_permissions"
	PermissionManageOrganizationUsers       Permission = "manage_organization_users"
	PermissionManageOrganizationRoles       Permission = "manage_organization_roles"
	PermissionManageOrganizationPermissions Permission = "manage_organization_permissions"
	PermissionViewOrganizationUsers         Permission = "view_organization_users"
	PermissionChangeOwnPassword             Permission = "change_own_password"
)

type DescriptionSystemRole string

const (
	DescriptionAdmin             DescriptionSystemRole = "admin"
	DescriptionOrganizationOwner DescriptionSystemRole = "organization_owner"
	DescriptionManager           DescriptionSystemRole = "manager"
	DescriptionUser              DescriptionSystemRole = "user"
)

type Operation string

const (
	Create Operation = "create"
	Read   Operation = "read"
	Update Operation = "update"
	Delete Operation = "delete"
)
