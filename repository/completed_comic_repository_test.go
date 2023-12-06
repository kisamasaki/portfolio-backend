package repository

import (
	"portfolio-backend/model"
	"portfolio-backend/testhelper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCompletedComic(t *testing.T) {
	db, sqlDB := testhelper.SetupDatabase(t)
	defer sqlDB.Close()

	userRepo := NewUserRepository(db)
	completedComicRepo := NewCompletedComicRepository(db)

	user := testhelper.TestUsers["testuser1"]
	userAuth := testhelper.TestUserAuths["testuser1"]
	if err := userRepo.CreateUser(&user, &userAuth); err != nil {
		t.Fatalf("Error creating user: %v", err)
	}

	t.Run("正常に読み終えた作品を登録することができる", func(t *testing.T) {
		err := completedComicRepo.CreateCompletedComic(&testhelper.TestCompletedComic)
		if err != nil {
			t.Fatalf("Error creating completed comic: %v", err)
		}
		assert.NoError(t, err)
	})
}

func TestCheckCompletedComicStatusFalse(t *testing.T) {
	db, sqlDB := testhelper.SetupDatabase(t)
	defer sqlDB.Close()

	userRepo := NewUserRepository(db)
	completedComicRepo := NewCompletedComicRepository(db)

	user := testhelper.TestUsers["testuser1"]
	userAuth := testhelper.TestUserAuths["testuser1"]
	if err := userRepo.CreateUser(&user, &userAuth); err != nil {
		t.Fatalf("Error creating user: %v", err)
	}

	t.Run("未読作品ならfalseを返す", func(t *testing.T) {
		isCompletedComicStatus, err := completedComicRepo.CheckCompletedComicStatus(user.ID, testhelper.TestComic.ID)
		assert.NoError(t, err)
		assert.Equal(t, false, isCompletedComicStatus)
	})
}

func TestCheckCompletedComicStatusTrue(t *testing.T) {
	db, sqlDB := testhelper.SetupDatabase(t)
	defer sqlDB.Close()

	userRepo := NewUserRepository(db)
	completedComicRepo := NewCompletedComicRepository(db)

	user := testhelper.TestUsers["testuser1"]
	userAuth := testhelper.TestUserAuths["testuser1"]
	if err := userRepo.CreateUser(&user, &userAuth); err != nil {
		t.Fatalf("Error creating user: %v", err)
	}

	if err := completedComicRepo.CreateCompletedComic(&testhelper.TestCompletedComic); err != nil {
		t.Fatalf("Error creating CompletedComic: %v", err)
	}

	t.Run("読み終えた作品ならtrueを返す", func(t *testing.T) {
		isCompletedComicStatus, err := completedComicRepo.CheckCompletedComicStatus(user.ID, testhelper.TestComic.ID)
		assert.NoError(t, err)
		assert.Equal(t, true, isCompletedComicStatus)
	})
}

func TestGetUserCompletedComics(t *testing.T) {
	db, sqlDB := testhelper.SetupDatabase(t)
	defer sqlDB.Close()

	userRepo := NewUserRepository(db)
	completedComicRepo := NewCompletedComicRepository(db)

	user := testhelper.TestUsers["testuser1"]
	userAuth := testhelper.TestUserAuths["testuser1"]
	if err := userRepo.CreateUser(&user, &userAuth); err != nil {
		t.Fatalf("Error creating user: %v", err)
	}

	if err := completedComicRepo.CreateCompletedComic(&testhelper.TestCompletedComic); err != nil {
		t.Fatalf("Error creating CompletedComic: %v", err)
	}

	t.Run("ユーザーが読み終えた作品を返す", func(t *testing.T) {
		completedComicRes := []model.CompletedComic{}
		err := completedComicRepo.GetUserCompletedComics(&completedComicRes, user.ID, 1)
		assert.NoError(t, err)
		assert.Equal(t, testhelper.TestCompletedComic.ID, completedComicRes[0].ID)
		assert.Equal(t, testhelper.TestCompletedComic.Rating, completedComicRes[0].Rating)
		assert.Equal(t, testhelper.TestCompletedComic.Review, completedComicRes[0].Review)
		assert.Equal(t, testhelper.TestCompletedComic.UserID, completedComicRes[0].UserID)
		assert.Equal(t, testhelper.TestCompletedComic.ComicID, completedComicRes[0].ComicID)
	})
}

func TestGetFolllowCompletedComics(t *testing.T) {
	db, sqlDB := testhelper.SetupDatabase(t)
	defer sqlDB.Close()

	userRepo := NewUserRepository(db)
	followingRepo := NewFollowingRepository(db)
	completedComicRepo := NewCompletedComicRepository(db)

	followUser := testhelper.TestUsers["testuser2"]
	followUserAuth := testhelper.TestUserAuths["testuser2"]
	followerUser := testhelper.TestUsers["testuser1"]
	followerUserAuth := testhelper.TestUserAuths["testuser1"]
	if err := userRepo.CreateUser(&followUser, &followUserAuth); err != nil {
		t.Fatalf("Error creating user: %v", err)
	}

	if err := userRepo.CreateUser(&followerUser, &followerUserAuth); err != nil {
		t.Fatalf("Error creating user: %v", err)
	}

	if err := followingRepo.Follow(followUser.ID, followerUser.ID); err != nil {
		t.Fatalf("Error Following: %v", err)
	}

	if err := completedComicRepo.CreateCompletedComic(&testhelper.TestCompletedComic); err != nil {
		t.Fatalf("Error creating CompletedComic: %v", err)
	}

	t.Run("フォローしているユーザーが読み終えた作品を返す", func(t *testing.T) {
		completedComicRes := []model.CompletedComic{}
		err := completedComicRepo.GetFolllowCompletedComics(&completedComicRes, followUser.ID, 1)
		assert.NoError(t, err)
		assert.Equal(t, testhelper.TestCompletedComic.ID, completedComicRes[0].ID)
		assert.Equal(t, testhelper.TestCompletedComic.Rating, completedComicRes[0].Rating)
		assert.Equal(t, testhelper.TestCompletedComic.Review, completedComicRes[0].Review)
		assert.Equal(t, testhelper.TestCompletedComic.UserID, completedComicRes[0].UserID)
		assert.Equal(t, testhelper.TestCompletedComic.ComicID, completedComicRes[0].ComicID)
	})
}

func TestGetLatestCompletedComics(t *testing.T) {
	db, sqlDB := testhelper.SetupDatabase(t)
	defer sqlDB.Close()

	userRepo := NewUserRepository(db)
	completedComicRepo := NewCompletedComicRepository(db)

	user := testhelper.TestUsers["testuser1"]
	userAuth := testhelper.TestUserAuths["testuser1"]
	if err := userRepo.CreateUser(&user, &userAuth); err != nil {
		t.Fatalf("Error creating user: %v", err)
	}

	if err := completedComicRepo.CreateCompletedComic(&testhelper.TestCompletedComic); err != nil {
		t.Fatalf("Error creating CompletedComic: %v", err)
	}

	t.Run("最新の読み終えた作品を返す", func(t *testing.T) {
		completedComicRes := []model.CompletedComic{}
		err := completedComicRepo.GetLatestCompletedComics(&completedComicRes)
		assert.NoError(t, err)
		assert.Equal(t, testhelper.TestCompletedComic.ID, completedComicRes[0].ID)
		assert.Equal(t, testhelper.TestCompletedComic.Rating, completedComicRes[0].Rating)
		assert.Equal(t, testhelper.TestCompletedComic.Review, completedComicRes[0].Review)
		assert.Equal(t, testhelper.TestCompletedComic.UserID, completedComicRes[0].UserID)
		assert.Equal(t, testhelper.TestCompletedComic.ComicID, completedComicRes[0].ComicID)
	})
}
