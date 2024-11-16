package usecases

import (
	"Bangseungjae/cockroach/cockroach/entities"
	"Bangseungjae/cockroach/cockroach/models"
	"Bangseungjae/cockroach/cockroach/repositories"
	"time"
)

type cockroachUsecaseImpl struct {
	cockroachRepository repositories.CockroachRepository
	cockroachMessage    repositories.CockroachMessaging
}

func NewCockroachUsecaseImpl(
	cockroachRepository repositories.CockroachRepository,
	cockroachMessage repositories.CockroachMessaging,
) *cockroachUsecaseImpl {
	return &cockroachUsecaseImpl{cockroachRepository: cockroachRepository, cockroachMessage: cockroachMessage}
}

func (u cockroachUsecaseImpl) CockroachDataProcessing(in *models.AddCockroachData) error {
	insertCockroachData := &entities.InsertCockroachDto{
		Amount: in.Amount,
	}

	if err := u.cockroachRepository.InsertCockroachData(insertCockroachData); err != nil {
		return err
	}
	pushCockroachData := &entities.CockroachPushNotificationDto{
		Title:        "Cockroach Detected 🪳 !!!",
		Amount:       in.Amount,
		ReportedTime: time.Now().Local().Format("2006-01-02 15:04:05"),
	}

	if err := u.cockroachMessage.PushNotification(pushCockroachData); err != nil {
		return err
	}
	return nil
}
