package auth_test

import (
	"context"
	"testing"

	"dbfhub.gitlab.yandexcloud.net/plano-dev/backend/plano-auth.git/pkg/auth"

	"github.com/google/uuid"
)

func TestContext(t *testing.T) {
	t.Parallel()

	backgroundCtx := context.Background()
	storedClaims := auth.FromContext(backgroundCtx)

	if storedClaims != nil {
		t.Errorf("expected claims to be nil, got %+v", *storedClaims)
	}

	claims := &auth.UserClaims{
		UUID: uuid.New(),
	}

	ctx := auth.Context(backgroundCtx, claims)

	storedClaims = auth.FromContext(ctx)
	if storedClaims == nil {
		t.Errorf("expected claims to be %+v, got nil", *claims)
	} else if storedClaims.UUID != claims.UUID {
		t.Errorf("expected claims to be %+v, got %+v", *claims, *storedClaims)
	}
}
