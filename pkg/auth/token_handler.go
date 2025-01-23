package auth

import (
	"time"

	"errors"
	"github.com/golang-jwt/jwt"
)

// ErrInvalidClaims возвращается когда описание пользователя некорректно.
var ErrInvalidClaims = errors.New("invalid claims")

// TokenHandler это интерфейс обработчика для работы с токеном.
type TokenHandler interface {
	Generate(claims *UserClaims) (string, error)
	Parse(tokenString string) (*UserClaims, error)
	TTL() time.Duration
}

// TokenConfig это структура конфигурации обработчика для работы с токеном.
type TokenConfig struct {
	Secret string
	TTL    time.Duration
}

// JWTHandler это обработчик для работы с токеном.
type JWTHandler struct {
	key []byte
	ttl time.Duration
}

// NewJWTHandler создает обработчик для работы с токеном.
func NewJWTHandler(config *TokenConfig) TokenHandler {
	return &JWTHandler{
		key: []byte(config.Secret),
		ttl: config.TTL,
	}
}

// Generate генерирует токен из описания пользователя.
func (j *JWTHandler) Generate(claims *UserClaims) (string, error) {
	if claims == nil {
		return "", ErrInvalidClaims
	}

	claims.ExpiresAt = time.Now().Add(j.ttl).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(j.key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Parse достаёт описание пользователя из токена.
func (j *JWTHandler) Parse(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, j.provideKey)
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrInvalidKey
}

// TTL возвращает время жизни токена.
func (j *JWTHandler) TTL() time.Duration {
	return j.ttl
}

func (j *JWTHandler) provideKey(_ *jwt.Token) (any, error) {
	return j.key, nil
}
