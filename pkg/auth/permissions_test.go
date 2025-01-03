package auth_test

import (
	"encoding/json"
	"strings"
	"testing"

	"dbfhub.gitlab.yandexcloud.net/plano-dev/backend/plano-auth.git/pkg/auth"
)

func TestPermissionsHas(t *testing.T) {
	p := auth.NewPermissions("test1", "test2", "test3")
	if !p.Has("test1") || !p.Has("test2") || !p.Has("test3") {
		t.Errorf("expected permissions to contain test1, test2 and test3, got: %v", p)
	}
	if p.Has("test4") {
		t.Errorf("expected permissions to not contain test4, got: %v", p)
	}

	if !p.HasAll("test1", "test2", "test3") {
		t.Errorf("expected permissions to contain test1, test2 and test3, got: %v", p)
	}
	if p.HasAll("test1", "test2", "test3", "test4") {
		t.Errorf("expected permissions to not contain test4, got: %v", p)
	}

	if !p.HasAny("test1", "test2", "test3") {
		t.Errorf("expected permissions to contain test1, test2 and test3, got: %v", p)
	}
	if !p.HasAny("test1", "test2", "test3", "test4") {
		t.Errorf("expected permissions to not contain test4, got: %v", p)
	}
	if p.HasAny("test4") {
		t.Errorf("expected permissions to not contain test4, got: %v", p)
	}
}

func TestPermissionsAdd(t *testing.T) {
	p := auth.NewPermissions()

	p.Add("test1", "test2", "test3")
	if !p.Has("test1") || !p.Has("test2") || !p.Has("test3") {
		t.Errorf("expected permissions to contain test1, test2 and test3, got: %v", p)
	}
}

func TestPermissionsMarshallJSON(t *testing.T) {
	p := make(auth.Permissions)

	p.Add("test1", "test2", "test3")

	data, err := json.Marshal(&p)
	if err != nil {
		t.Errorf("failed to marshal permissions: %s", err.Error())
		return
	}

	jsonValue := string(data)
	if !strings.HasPrefix(jsonValue, "[") || !strings.HasSuffix(jsonValue, "]") {
		t.Errorf("expected permissions to be an array, got: %s", jsonValue)
		return
	}

	if !strings.Contains(jsonValue, "\"test1\"") ||
		!strings.Contains(jsonValue, "\"test2\"") ||
		!strings.Contains(jsonValue, "\"test3\"") {
		t.Errorf("expected permissions to contain test1, test2 and test3, got: %s", jsonValue)
	}
}

func TestPermissionsUnmarshallJSON(t *testing.T) {
	var p auth.Permissions
	err := json.Unmarshal([]byte(`["test1", "test2", "test3"]`), &p)
	if err != nil {
		t.Errorf("failed to unmarshal permissions: %s", err.Error())
		return
	}

	if !p.Has("test1") {
		t.Errorf("expected permissions to contain test1, got: %v", p)
	}
	if !p.Has("test2") {
		t.Errorf("expected permissions to contain test2, got: %v", p)
	}
	if !p.Has("test3") {
		t.Errorf("expected permissions to contain test3, got: %v", p)
	}
}
