package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"portfolio-backend/model"
	"portfolio-backend/repository"
	"portfolio-backend/testhelper"
	"portfolio-backend/usecase"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestFollow(t *testing.T) {

	db, sqlDB := testhelper.SetupDatabase(t)
	defer sqlDB.Close()

	userRepo := repository.NewUserRepository(db)
	followingRepo := repository.NewFollowingRepository(db)
	followingUsecase := usecase.NewFollowingUsecase(followingRepo)
	followingController := NewFollowingController(followingUsecase)

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

	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/follow/:userId/:followerId")
	c.SetParamNames("userId", "followerId")
	c.SetParamValues(followUser.ID, followerUser.ID)

	t.Run("正常にフォローすることができる", func(t *testing.T) {
		err := followingController.Follow(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

func TestUnfollow(t *testing.T) {

	db, sqlDB := testhelper.SetupDatabase(t)
	defer sqlDB.Close()

	userRepo := repository.NewUserRepository(db)
	followingRepo := repository.NewFollowingRepository(db)
	followingUsecase := usecase.NewFollowingUsecase(followingRepo)
	followingController := NewFollowingController(followingUsecase)

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

	if err := followingRepo.Unfollow(followUser.ID, followerUser.ID); err != nil {
		t.Fatalf("Error Unfollowing: %v", err)
	}

	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/follow/unfollow/:userId/:unfollowUserId")
	c.SetParamNames("userId", "unfollowUserId")
	c.SetParamValues(followUser.ID, followerUser.ID)

	t.Run("正常にフォローを解除することができる", func(t *testing.T) {
		err := followingController.Unfollow(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

func TestCheckFollowingStatusFalse(t *testing.T) {
	db, sqlDB := testhelper.SetupDatabase(t)
	defer sqlDB.Close()

	userRepo := repository.NewUserRepository(db)
	followingRepo := repository.NewFollowingRepository(db)
	followingUsecase := usecase.NewFollowingUsecase(followingRepo)
	followingController := NewFollowingController(followingUsecase)

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

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/follow/check-follow/:sessionUserId/:userId")
	c.SetParamNames("sessionUserId", "userId")
	c.SetParamValues(followUser.ID, followerUser.ID)

	t.Run("未フォロー状態の場合falseを返す", func(t *testing.T) {
		err := followingController.CheckFollowingStatus(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		response := map[string]interface{}{}
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		isFollowing := response["isFollowing"].(bool)
		assert.False(t, isFollowing)
	})
}

func TestCheckFollowingStatusTrue(t *testing.T) {
	db, sqlDB := testhelper.SetupDatabase(t)
	defer sqlDB.Close()

	userRepo := repository.NewUserRepository(db)
	followingRepo := repository.NewFollowingRepository(db)
	followingUsecase := usecase.NewFollowingUsecase(followingRepo)
	followingController := NewFollowingController(followingUsecase)

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
		t.Fatalf("Error creating Following: %v", err)
	}

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/follow/check-follow/:sessionUserId/:userId")
	c.SetParamNames("sessionUserId", "userId")
	c.SetParamValues(followUser.ID, followerUser.ID)

	t.Run("フォロー状態の場合trueを返す", func(t *testing.T) {
		err := followingController.CheckFollowingStatus(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		response := map[string]interface{}{}
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		isFollowing := response["isFollowing"].(bool)
		assert.True(t, isFollowing)
	})
}

func TestGetFollowingUsers(t *testing.T) {
	db, sqlDB := testhelper.SetupDatabase(t)
	defer sqlDB.Close()

	userRepo := repository.NewUserRepository(db)
	followingRepo := repository.NewFollowingRepository(db)
	followingUsecase := usecase.NewFollowingUsecase(followingRepo)
	followingController := NewFollowingController(followingUsecase)

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
		t.Fatalf("Error creating Following: %v", err)
	}

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/follow/following/:userId")
	c.SetParamNames("userId")
	c.SetParamValues(followUser.ID)

	t.Run("正常にフォローユーザーを取得することができる", func(t *testing.T) {
		err := followingController.GetFollowingUsers(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		var followerUserRes []model.User
		err = json.Unmarshal(rec.Body.Bytes(), &followerUserRes)
		assert.NoError(t, err)
		assert.Equal(t, followerUser.ID, followerUserRes[0].ID)
		assert.Equal(t, followerUser.UserName, followerUserRes[0].UserName)
	})
}

func TestGetFollowerUsers(t *testing.T) {
	db, sqlDB := testhelper.SetupDatabase(t)
	defer sqlDB.Close()

	userRepo := repository.NewUserRepository(db)
	followingRepo := repository.NewFollowingRepository(db)
	followingUsecase := usecase.NewFollowingUsecase(followingRepo)
	followingController := NewFollowingController(followingUsecase)

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

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/follow/followers/:userId")
	c.SetParamNames("userId")
	c.SetParamValues(followerUser.ID)

	t.Run("正常にフォロワーを取得することができる", func(t *testing.T) {
		err := followingController.GetFollowerUsers(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		var followingUserRes []model.User
		err = json.Unmarshal(rec.Body.Bytes(), &followingUserRes)
		assert.NoError(t, err)
		assert.Equal(t, followUser.ID, followingUserRes[0].ID)
		assert.Equal(t, followUser.UserName, followingUserRes[0].UserName)
	})
}
