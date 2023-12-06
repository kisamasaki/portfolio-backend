package repository

import (
	"portfolio-backend/testhelper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateComics_Normal(t *testing.T) {
	db, sqlDB := testhelper.SetupDatabase(t)
	defer sqlDB.Close()

	repo := NewComicRepository(db, nil)
	comics := testhelper.TestComics
	savedComics := testhelper.TestComics

	t.Run("正常に登録することができる", func(t *testing.T) {
		err := repo.CreateComics(&comics)
		assert.NoError(t, err)
		db.Find(&savedComics)
		assert.Equal(t, len(comics), len(savedComics))
	})
}

func TestCreateComics_Duplicate(t *testing.T) {
	db, sqlDB := testhelper.SetupDatabase(t)
	defer sqlDB.Close()

	repo := NewComicRepository(db, nil)
	comics := testhelper.TestComics
	duplicateComics := testhelper.TestComics

	t.Run("重複があった場合、エラーが発生しない", func(t *testing.T) {
		if err := repo.CreateComics(&comics); err != nil {
			t.Fatalf("Error creating comics: %v", err)
		}
		err := repo.CreateComics(&duplicateComics)
		assert.NoError(t, err)
	})
}
