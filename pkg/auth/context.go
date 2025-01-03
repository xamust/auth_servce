package auth

import "context"

type contextKey string

const claimsKey contextKey = "authUserClaims"

// Context возвращает обновлённый контекст с данными о пользователе
func Context(ctx context.Context, claims *UserClaims) context.Context {
	return context.WithValue(ctx, claimsKey, claims)
}

// FromContext возвращает данные о пользователе из контекста
func FromContext(ctx context.Context) *UserClaims {
	if claims, ok := ctx.Value(claimsKey).(*UserClaims); ok {
		return claims
	}
	return nil
}
