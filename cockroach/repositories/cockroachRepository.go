package repositories

import "Bangseungjae/cockroach/cockroach/entities"

type CockroachRepository interface {
	InsertCockroachData(in *entities.InsertCockroachDto) error
}
