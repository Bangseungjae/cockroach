package tests

import (
	"Bangseungjae/cockroach/cockroach/entities"
	"Bangseungjae/cockroach/config"
	"Bangseungjae/cockroach/database"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"strings"
	"sync"
	"testing"
)

var (
	once           sync.Once
	configInstance *config.Config
)

func GetConfig() *config.Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}
	})

	return configInstance
}

func truncateAllTables(db *gorm.DB) error {
	var tables []string

	// PostgreSQL의 public 스키마에 있는 모든 테이블 이름을 조회합니다.
	if err := db.Raw("SELECT tablename FROM pg_tables WHERE schemaname = 'public'").Scan(&tables).Error; err != nil {
		return err
	}

	if len(tables) == 0 {
		return nil // 트렁케이트할 테이블이 없음
	}

	// 테이블 이름을 쉼표로 구분된 문자열로 변환합니다.
	truncateStmt := "TRUNCATE TABLE " + strings.Join(tables, ", ") + " RESTART IDENTITY CASCADE;"

	// TRUNCATE 명령어 실행
	return db.Exec(truncateStmt).Error
}

func setupTestDB(t *testing.T) database.Database {
	// 테스트용 설정 로드
	conf := GetConfig()

	// 데이터베이스 초기화
	db := database.NewPostgresDatabase(conf)
	require.NotNil(t, db, "Database should be initialized")

	// 마이그레이션 실행
	err := db.GetDb().AutoMigrate(&entities.Cockroach{})
	require.NoError(t, err, "Database migration should not fail")

	// 트렁케이트 실행
	err = truncateAllTables(db.GetDb())
	require.NoError(t, err, "Failed to truncate tables before test")

	return db
}
