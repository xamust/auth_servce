package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Common struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
