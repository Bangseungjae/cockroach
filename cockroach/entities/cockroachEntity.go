package entities

import (
	"gorm.io/gorm"
	"time"
)

type (
	InsertCockroachDto struct {
		Id        uint32    `json:"id"`
		Amount    uint32    `json:"amount"`
		CreatedAt time.Time `json:"createdAt"`
	}

	Cockroach struct {
		Amount uint32 `json:"amount"`
		gorm.Model
	}

	CockroachPushNotificationDto struct {
		Title        string `json:"title"`
		Amount       uint32 `json:"amount"`
		ReportedTime string `json:"createdAt"`
	}
)
