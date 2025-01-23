package repositories

import (
	"context"
	"gitlab.com/xamops/auth/internal/interfaces"
	"gorm.io/gorm"
	"log/slog"
)

func New(db *gorm.DB, log *slog.Logger) *Repository {
	return &Repository{
		db:   db,
		auth: newAuthRepository(log),
		user: newUsersRepository(log),
		role: newSystemRolesRepository(log),
	}
}

type Repository struct {
	db   *gorm.DB
	auth interfaces.AuthRepository
	user interfaces.UsersRepository
	role interfaces.SystemRolesRepository
}

func (r *Repository) ConnWithContext(ctx context.Context) *gorm.DB {
	return r.Conn().WithContext(ctx)
}

func (r *Repository) Conn() *gorm.DB {
	return r.db
}

func (r *Repository) Auth() interfaces.AuthRepository {
	return r.auth
}

func (r *Repository) Users() interfaces.UsersRepository {
	return r.user
}

func (r *Repository) Roles() interfaces.SystemRolesRepository {
	return r.role
}
