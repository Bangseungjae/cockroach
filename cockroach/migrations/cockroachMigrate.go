package main

import (
	"Bangseungjae/cockroach/cockroach/entities"
	"Bangseungjae/cockroach/config"
	"Bangseungjae/cockroach/database"
	"fmt"
)

func main() {
	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)
	cockroachMigrate(db)
}

func cockroachMigrate(db database.Database) {
	db.GetDb().Migrator().DropTable(&entities.Cockroach{})
	db.GetDb().Migrator().CreateTable(&entities.Cockroach{})
	db.GetDb().CreateInBatches([]entities.Cockroach{
		{Amount: 1},
		{Amount: 2},
		{Amount: 2},
		{Amount: 5},
		{Amount: 3},
	}, 10)

	var cockroach entities.Cockroach
	db.GetDb().First(&cockroach, 1)
	fmt.Printf("Inserted Record: %+v\n", cockroach)
	// 소프트 삭제 수행
	db.GetDb().Delete(&cockroach)

	// 삭제된 레코드 조회 (기본적으로 제외됨)
	var deletedCockroach entities.Cockroach
	result := db.GetDb().First(&deletedCockroach, cockroach.ID)
	if result.Error != nil {
		fmt.Println("Record not found (soft deleted)")
	}

	// 삭제된 레코드도 포함하여 조회
	db.GetDb().Unscoped().First(&deletedCockroach, cockroach.ID)
	fmt.Printf("Soft Deleted Record: %+v\n", deletedCockroach)
}
