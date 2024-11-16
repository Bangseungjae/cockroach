package repositories

import (
	"Bangseungjae/cockroach/cockroach/entities"
	"Bangseungjae/cockroach/database"
	"context"
	"gorm.io/gorm"
)

type CockroachRepository interface {
	InsertCockroachData(ctx context.Context, tx *gorm.DB, in *entities.InsertCockroachDto) error
	GetDB() database.Database
}
