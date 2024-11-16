package entities

import (
	"database/sql"
	"time"
)

type DeletedAt sql.NullTime

type CommonEntityBase struct {
	Id        uint32 `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt DeletedAt `gorm:"index"`
}
