package tests

import (
	"Bangseungjae/cockroach/cockroach/entities"
	"Bangseungjae/cockroach/cockroach/models"
	"Bangseungjae/cockroach/cockroach/repositories"
	"Bangseungjae/cockroach/cockroach/usecases"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCockroachDataProcessing(t *testing.T) {
	// Step 1: 시작

	// Step 2: 테스트 데이터베이스 초기화
	db := setupTestDB(t)
	t.Cleanup(func() {
		truncateAllTables(db.GetDb())
	})

	require.NotNil(t, db, "Database should be initialized")

	require.NoError(t, db.GetDb().Error, "Failed to begin transaction")

	// Step 4: 테스트 전에 기존 데이터 삽입 (선택 사항)
	err := db.GetDb().Where("1 = 1").Delete(&entities.Cockroach{}).Error
	require.NoError(t, err, "Failed to clean up database before test")

	// Step 5: 실제 구현체를 사용하여 리포지토리 초기화
	cockroachRepository := repositories.NewCockroachPostgresRepository(db)
	require.NotNil(t, cockroachRepository, "CockroachRepository should be initialized")

	cockroachMessaging := repositories.NewCockroachFCMMessaging() // 실제 구현된 생성자 사용
	require.NotNil(t, cockroachMessaging, "CockroachMessaging should be initialized")

	// Step 6: 실제 의존성을 사용하여 유스케이스 초기화
	usecase := usecases.NewCockroachUsecaseImpl(cockroachRepository, cockroachMessaging)
	require.NotNil(t, usecase, "CockroachUsecase should be initialized")

	// Step 7: 테스트 입력 데이터 준비
	testInput := &models.AddCockroachData{
		Amount: 10,
	}

	// Step 8: 유즈케이스 메서드 실행
	ctx := context.Background()

	err = usecase.CockroachDataProcessing(ctx, testInput)
	assert.NoError(t, err, "CockroachDataProcessing should not return an error")

	// Step 9: 데이터베이스에 데이터가 삽입되었는지 확인
	var insertedData entities.Cockroach
	err = db.GetDb().Where("amount = ?", testInput.Amount).First(&insertedData).Error
	assert.NoError(t, err, "Inserted Cockroach data should exist in the database")
	assert.Equal(t, testInput.Amount, insertedData.Amount, "Amount should match the input")

	// Step 10: CreatedAt 필드가 올바르게 설정되었는지 확인
	assert.False(t, insertedData.CreatedAt.IsZero(), "CreatedAt should be set")
	assert.WithinDuration(t, time.Now(), insertedData.CreatedAt, time.Second*5, "CreatedAt should be recent")

	// Step 11: 푸시 알림이 제대로 전송되었는지 확인 (선택 사항)
	// 실제 FCM 서비스와 연동되어 있다면, 메시지가 전송되었는지 확인할 방법이 필요합니다.
	// 예를 들어, FCM 로그를 확인하거나 테스트용 FCM 토큰을 사용하여 수신 여부를 확인할 수 있습니다.
	// 여기서는 단순히 로그를 통해 확인할 수 있도록 메시징 구현체가 로그를 남긴다고 가정합니다.

}
