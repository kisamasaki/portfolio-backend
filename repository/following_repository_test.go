package repository

import (
	"portfolio-backend/testhelper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFollow(t *testing.T) {
	db, sqlDB := testhelper.SetupDatabase(t)
	defer sqlDB.Close()

	userRepo := NewUserRepository(db)
	followingRepo := NewFollowingRepository(db)

	followUser := testhelper.TestUsers["testuser1"]
	followUserAuth := testhelper.TestUserAuths["testuser1"]
	followerUser := testhelper.TestUsers["testuser2"]
	followerUserAuth := testhelper.TestUserAuths["testuser2"]

	if err := userRepo.CreateUser(&followUser, &followUserAuth); err != nil {
		t.Fatalf("Error creating user: %v", err)
	}

	if err := userRepo.CreateUser(&followerUser, &followerUserAuth); err != nil {
		t.Fatalf("Error creating user: %v", err)
	}

	t.Run("正常にフォローすることができる", func(t *testing.T) {
		err := followingRepo.Follow(followUser.ID, followerUser.ID)
		assert.NoError(t, err)
	})
}

func TestUnfollow(t *testing.T) {
	db, sqlDB := testhelper.SetupDatabase(t)
	defer sqlDB.Close()

	userRepo := NewUserRepository(db)
	followingRepo := NewFollowingRepository(db)

	followUser := testhelper.TestUsers["testuser1"]
	followUserAuth := testhelper.TestUserAuths["testuser1"]
	followerUser := testhelper.TestUsers["testuser2"]
	followerUserAuth := testhelper.TestUserAuths["testuser2"]

	if err := userRepo.CreateUser(&followUser, &followUserAuth); err != nil {
		t.Fatalf("Error creating user: %v", err)
	}

	if err := userRepo.CreateUser(&followerUser, &followerUserAuth); err != nil {
		t.Fatalf("Error creating user: %v", err)
	}

	if err := followingRepo.Follow(followUser.ID, followerUser.ID); err != nil {
		t.Fatalf("Error Following: %v", err)
	}

	t.Run("正常にフォローを解除することができる", func(t *testing.T) {
		err := followingRepo.Unfollow(followUser.ID, followerUser.ID)
		assert.NoError(t, err)
	})
}

func TestCheckFollowingStatusFalse(t *testing.T) {
	db, sqlDB := testhelper.SetupDatabase(t)
	defer sqlDB.Close()

	userRepo := NewUserRepository(db)
	followingRepo := NewFollowingRepository(db)

	followUser := testhelper.TestUsers["testuser1"]
	followUserAuth := testhelper.TestUserAuths["testuser1"]
	followerUser := testhelper.TestUsers["testuser2"]
	followerUserAuth := testhelper.TestUserAuths["testuser2"]

	if err := userRepo.CreateUser(&followUser, &followUserAuth); err != nil {
		t.Fatalf("Error creating user: %v", err)
	}

	if err := userRepo.CreateUser(&followerUser, &followerUserAuth); err != nil {
		t.Fatalf("Error creating user: %v", err)
	}

	t.Run("未フォロー状態の場合falseを返す", func(t *testing.T) {
		isFollowingStatus, err := followingRepo.CheckFollowingStatus(followUser.ID, followerUser.ID)
		assert.NoError(t, err)
		assert.Equal(t, false, isFollowingStatus)
	})
}

func TestCheckFollowingStatusTrue(t *testing.T) {
	db, sqlDB := testhelper.SetupDatabase(t)
	defer sqlDB.Close()

	userRepo := NewUserRepository(db)
	followingRepo := NewFollowingRepository(db)

	followUser := testhelper.TestUsers["testuser1"]
	followUserAuth := testhelper.TestUserAuths["testuser1"]
	followerUser := testhelper.TestUsers["testuser2"]
	followerUserAuth := testhelper.TestUserAuths["testuser2"]

	if err := userRepo.CreateUser(&followUser, &followUserAuth); err != nil {
		t.Fatalf("Error creating user: %v", err)
	}

	if err := userRepo.CreateUser(&followerUser, &followerUserAuth); err != nil {
		t.Fatalf("Error creating user: %v", err)
	}

	if err := followingRepo.Follow(followUser.ID, followerUser.ID); err != nil {
		t.Fatalf("Error Following: %v", err)
	}

	t.Run("フォロー状態の場合trueを返す", func(t *testing.T) {
		isFollowingStatus, err := followingRepo.CheckFollowingStatus(followUser.ID, followerUser.ID)
		assert.NoError(t, err)
		assert.Equal(t, true, isFollowingStatus)
	})
}

func TestGetFollowingUsers(t *testing.T) {
	db, sqlDB := testhelper.SetupDatabase(t)
	defer sqlDB.Close()

	userRepo := NewUserRepository(db)
	followingRepo := NewFollowingRepository(db)

	followUser := testhelper.TestUsers["testuser1"]
	followUserAuth := testhelper.TestUserAuths["testuser1"]
	followerUser := testhelper.TestUsers["testuser2"]
	followerUserAuth := testhelper.TestUserAuths["testuser2"]

	if err := userRepo.CreateUser(&followUser, &followUserAuth); err != nil {
		t.Fatalf("Error creating user: %v", err)
	}

	if err := userRepo.CreateUser(&followerUser, &followerUserAuth); err != nil {
		t.Fatalf("Error creating user: %v", err)
	}

	if err := followingRepo.Follow(followUser.ID, followerUser.ID); err != nil {
		t.Fatalf("Error Following: %v", err)
	}

	t.Run("正常にフォローユーザーを取得することができる", func(t *testing.T) {
		followerUserRes, err := followingRepo.GetFollowingUsers(followUser.ID)
		assert.NoError(t, err)
		assert.Equal(t, followerUser.ID, followerUserRes[0].ID)
		assert.Equal(t, followerUser.UserName, followerUserRes[0].UserName)
	})
}

func TestGetFollowerUsers(t *testing.T) {
	db, sqlDB := testhelper.SetupDatabase(t)
	defer sqlDB.Close()

	userRepo := NewUserRepository(db)
	followingRepo := NewFollowingRepository(db)

	followUser := testhelper.TestUsers["testuser1"]
	followUserAuth := testhelper.TestUserAuths["testuser1"]
	followerUser := testhelper.TestUsers["testuser2"]
	followerUserAuth := testhelper.TestUserAuths["testuser2"]

	if err := userRepo.CreateUser(&followUser, &followUserAuth); err != nil {
		t.Fatalf("Error creating user: %v", err)
	}

	if err := userRepo.CreateUser(&followerUser, &followerUserAuth); err != nil {
		t.Fatalf("Error creating user: %v", err)
	}

	if err := followingRepo.Follow(followUser.ID, followerUser.ID); err != nil {
		t.Fatalf("Error Following: %v", err)
	}

	t.Run("正常にフォロワーを取得することができる", func(t *testing.T) {
		followingUserRes, err := followingRepo.GetFollowerUsers(followerUser.ID)
		assert.NoError(t, err)
		assert.Equal(t, followUser.ID, followingUserRes[0].ID)
		assert.Equal(t, followUser.UserName, followingUserRes[0].UserName)
	})
}
