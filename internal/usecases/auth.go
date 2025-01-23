package usecases

import (
	"context"
	"fmt"
	"gitlab.com/xamops/auth/internal/config"
	"gitlab.com/xamops/auth/internal/entity"
	"gitlab.com/xamops/auth/internal/interfaces"
	//"gitlab.com/xamops/auth/pkg/auth"
	"golang.org/x/crypto/bcrypt"
	//"google.golang.org/grpc/codes"
	//"google.golang.org/grpc/status"
	"log/slog"
)

type AuthUsecases struct {
	config  *config.Config
	repo    interfaces.Repositories
	mappers interfaces.Mappers
	log     *slog.Logger
}

var _ interfaces.AuthUsecases = (*AuthUsecases)(nil)

func newAuthUsecases(deps Dependencies) *AuthUsecases {
	return &AuthUsecases{
		config:  deps.Config,
		repo:    deps.Repositories,
		mappers: deps.Mappers,
		log:     deps.Logger,
	}
}

func (a *AuthUsecases) ChangePassword(ctx context.Context, email string, passwordOld string, passwordNew string) error {
	var hash string
	var err error
	hash, err = a.repo.Auth().PasswordHashForUser(a.repo.ConnWithContext(ctx), email)
	if err != nil {
		//if errors.Is(err, gorm.ErrRecordNotFound) {
		//	return errors.NotFound
		//}
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwordOld))
	if err != nil {
		//return errors.Unauthorized
		return fmt.Errorf("wrong password")
	}

	newhash, err := bcrypt.GenerateFromPassword([]byte(passwordNew), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	a.log.Debug("auth.ChangePassword", "hash", string(newhash))

	err = a.repo.Auth().ChangePassword(a.repo.ConnWithContext(ctx), email, string(newhash))
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthUsecases) Login(ctx context.Context, email string, password string) (*entity.User, error) {
	const op = "AuthUsecase.Login"
	var hash string
	var err error
	hash, err = a.repo.Auth().PasswordHashForUser(a.repo.ConnWithContext(ctx), email)
	if err != nil {
		// Do not disclose user availability
		//if errors.Is(err, errors.NotFound) {
		//	return nil, errors.Unauthorized
		//}
		a.log.Error(op, "error", err.Error())
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("unauth")
	}

	user, err := a.repo.Users().ByEmail(a.repo.ConnWithContext(ctx), email)
	if err != nil {
		a.log.Error(op, "error", err.Error())
		return nil, err
	}

	return user, nil
}
