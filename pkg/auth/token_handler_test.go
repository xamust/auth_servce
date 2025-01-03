package auth_test

import (
	"testing"
	"time"

	"dbfhub.gitlab.yandexcloud.net/plano-dev/backend/plano-auth.git/pkg/auth"

	"github.com/google/uuid"
)

func TestJWTHandlerGenerate(t *testing.T) {
	t.Parallel()
	handler := auth.NewJWTHandler(&auth.TokenConfig{
		Secret: "secret",
		TTL:    1,
	})
	_, err := handler.Generate(&auth.UserClaims{
		UUID: uuid.New(),
	})
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}

	_, err = handler.Generate(nil)
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestJWTHandlerParse(t *testing.T) {
	t.Parallel()
	handler := auth.NewJWTHandler(&auth.TokenConfig{
		Secret: "secret",
		TTL:    1,
	})

	userID := uuid.New()
	token, err := handler.Generate(&auth.UserClaims{
		UUID: userID,
	})
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}

	claims, err := handler.Parse(token)
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	} else if claims == nil {
		t.Error("expected claims to be not nil")
	} else if claims.UUID != userID {
		t.Errorf("expected claims to be %+v, got %+v", userID, claims.UUID)
	}

	_, err = handler.Parse("")
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestJWTHandlerTTL(t *testing.T) {
	t.Parallel()

	ttl := time.Second * 5
	expiresAt := time.Now().Add(ttl).Unix()

	handler := auth.NewJWTHandler(&auth.TokenConfig{
		Secret: "secret",
		TTL:    ttl,
	})

	if handler.TTL() != ttl {
		t.Errorf("expected ttl to be %v, got %v", ttl, handler.TTL())
	}

	token, _ := handler.Generate(&auth.UserClaims{})
	claims, _ := handler.Parse(token)
	if claims.ExpiresAt != expiresAt {
		t.Errorf("expected claims to be %+v, got %+v", expiresAt, claims.ExpiresAt)
	}
}
