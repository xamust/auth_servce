package auth_test

import (
	"testing"

	"dbfhub.gitlab.yandexcloud.net/plano-dev/backend/plano-auth.git/pkg/auth"
)

func TestOperationPermissionsAllowed(t *testing.T) {
	const (
		writePermission auth.Permission = "write"
		readPermission  auth.Permission = "read"
	)

	alicePermissions := auth.NewPermissions(writePermission, readPermission)
	bobPermissions := auth.NewPermissions(readPermission)

	operationPermissions := auth.OperationPermissions[string]{
		"Read":  readPermission,
		"Write": writePermission,
	}

	if !operationPermissions.Allowed("Read", alicePermissions) {
		t.Errorf("expected alice to be allowed to read, got: %v", alicePermissions)
	}
	if !operationPermissions.Allowed("Write", alicePermissions) {
		t.Errorf("expected alice to be allowed to write, got: %v", alicePermissions)
	}
	if operationPermissions.Allowed("Write", bobPermissions) {
		t.Errorf("expected bob to be not allowed to write, got: %v", bobPermissions)
	}

}
