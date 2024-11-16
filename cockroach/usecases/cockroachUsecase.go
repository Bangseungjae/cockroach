package usecases

import "Bangseungjae/cockroach/cockroach/models"

type CockroachUsecase interface {
	CockroachDataProcessing(in *models.AddCockroachData) error
}
