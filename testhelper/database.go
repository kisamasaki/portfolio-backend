package testhelper

import (
	"database/sql"
	"portfolio-backend/model"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase(t *testing.T) (*gorm.DB, *sql.DB) {
	url := "host=localhost port=5435 user=testuser password=testpassword dbname=testdb sslmode=disable"
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		t.Fatalf("接続に失敗: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("DBの取得に失敗: %v", err)
	}

	// トランザクションを開始
	tx := db.Begin()
	if tx.Error != nil {
		t.Fatalf("トランザクションの開始に失敗: %v", tx.Error)
	}
	defer func() {
		// テスト終了時にトランザクションをロールバック
		tx.Rollback()
	}()

	db.Where("1=1").Delete(&model.CompletedComic{})
	db.Where("1=1").Delete(&model.Following{})
	db.Where("1=1").Delete(&model.Comic{})
	db.Where("1=1").Delete(&model.UserAuth{})
	db.Where("1=1").Delete(&model.User{})

	// トランザクションをコミット
	if err := tx.Commit().Error; err != nil {
		t.Fatalf("トランザクションのコミットに失敗: %v", err)
	}

	return db, sqlDB
}
