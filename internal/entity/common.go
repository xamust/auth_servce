package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Common struct {
	UUID      uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
