package repositories

import (
	"Bangseungjae/cockroach/cockroach/entities"
	"github.com/labstack/gommon/log"
)

type cockroachFCMMessaging struct{}

func NewCockroachFCMMessaging() *cockroachFCMMessaging {
	return &cockroachFCMMessaging{}
}

func (c cockroachFCMMessaging) PushNotification(m *entities.CockroachPushNotificationDto) error {
	// ... handle logic to push FCM notification here ...
	log.Debugf("Push FCM notification with data: %v", m)
	return nil
}
