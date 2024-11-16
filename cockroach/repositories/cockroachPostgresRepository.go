package repositories

import (
	"Bangseungjae/cockroach/cockroach/entities"
	"Bangseungjae/cockroach/database"
	"context"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type cockroachPostgresRepository struct {
	db database.Database
}

func (r cockroachPostgresRepository) GetDB() database.Database {
	return r.db
}

func NewCockroachPostgresRepository(db database.Database) *cockroachPostgresRepository {
	return &cockroachPostgresRepository{db: db}
}

func (r cockroachPostgresRepository) InsertCockroachData(ctx context.Context, tx *gorm.DB, in *entities.InsertCockroachDto) error {
	data := &entities.Cockroach{
		Amount: in.Amount,
	}

	result := tx.WithContext(ctx).Create(data)

	if result.Error != nil {
		log.Errorf("InsertCockroachData: %v", result.Error)
		return result.Error
	}

	log.Debugf("InsertCockroachData: %v", result.RowsAffected)
	return nil
}
