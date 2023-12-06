package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"portfolio-backend/model"
	"portfolio-backend/repository"
	"portfolio-backend/testhelper"
	"portfolio-backend/usecase"
	"portfolio-backend/validator"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateCompletedComic(t *testing.T) {
	db, sqlDB := testhelper.SetupDatabase(t)
	defer sqlDB.Close()

	comicRepo := repository.NewComicRepository(db, nil)
	userRepo := repository.NewUserRepository(db)
	completedComicRepo := repository.NewCompletedComicRepository(db)
	ccv := validator.NewCompletedComicValidator()
	completedComicUsecase := usecase.NewCompletedComicUsecase(completedComicRepo, ccv)
	completedComicController := NewCompletedComicController(completedComicUsecase)

	comics := []model.Comic{testhelper.TestComic}
	if err := comicRepo.CreateComics(&comics); err != nil {
		t.Fatalf("Error creating comics: %v", err)
	}

	user := testhelper.TestUsers["testuser1"]
	userAuth := testhelper.TestUserAuths["testuser1"]
	if err := userRepo.CreateUser(&user, &userAuth); err != nil {
		t.Fatalf("Error creating user: %v", err)
	}

	e := echo.New()
	requestBody := `{"itemNumber": "123", "rating": 5, "review": "記録用"}`
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/completedComics/:userId")
	c.SetParamNames("userId")
	c.SetParamValues(user.ID)

	t.Run("正常に読み終えた作品を登録することができる", func(t *testing.T) {
		err := completedComicController.CreateCompletedComic(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)
	})
}

func TestCheckCompletedComicStatusFalse(t *testing.T) {
	db, sqlDB := testhelper.SetupDatabase(t)
	defer sqlDB.Close()

	userRepo := repository.NewUserRepository(db)
	completedComicRepo := repository.NewCompletedComicRepository(db)
	ccv := validator.NewCompletedComicValidator()
	completedComicUsecase := usecase.NewCompletedComicUsecase(completedComicRepo, ccv)
	completedComicController := NewCompletedComicController(completedComicUsecase)

	user := testhelper.TestUsers["testuser1"]
	userAuth := testhelper.TestUserAuths["testuser1"]
	if err := userRepo.CreateUser(&user, &userAuth); err != nil {
		t.Fatalf("Error creating user: %v", err)
	}

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/completedComics/check-completedcomic/:comicId/:userId")
	c.SetParamNames("comicId", "userId")
	c.SetParamValues(testhelper.TestComic.ID, user.ID)

	t.Run("未読作品ならfalseを返す", func(t *testing.T) {
		err := completedComicController.CheckCompletedComicStatus(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		var response bool
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.False(t, response)
	})
}

func TestCheckCompletedComicStatusTrue(t *testing.T) {
	db, sqlDB := testhelper.SetupDatabase(t)
	defer sqlDB.Close()

	userRepo := repository.NewUserRepository(db)
	completedComicRepo := repository.NewCompletedComicRepository(db)
	ccv := validator.NewCompletedComicValidator()
	completedComicUsecase := usecase.NewCompletedComicUsecase(completedComicRepo, ccv)
	completedComicController := NewCompletedComicController(completedComicUsecase)

	user := testhelper.TestUsers["testuser1"]
	userAuth := testhelper.TestUserAuths["testuser1"]
	if err := userRepo.CreateUser(&user, &userAuth); err != nil {
		t.Fatalf("Error creating user: %v", err)
	}
	if err := completedComicRepo.CreateCompletedComic(&testhelper.TestCompletedComic); err != nil {
		t.Fatalf("Error creating CreateCompletedComic: %v", err)
	}

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/completedComics/check-completedcomic/:comicId/:userId")
	c.SetParamNames("comicId", "userId")
	c.SetParamValues(testhelper.TestComic.ID, user.ID)

	t.Run("読み終えた作品ならtrueを返す", func(t *testing.T) {
		err := completedComicController.CheckCompletedComicStatus(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		var response bool
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.True(t, response)
	})
}

func TestGetUserCompletedComics(t *testing.T) {
	db, sqlDB := testhelper.SetupDatabase(t)
	defer sqlDB.Close()

	userRepo := repository.NewUserRepository(db)
	completedComicRepo := repository.NewCompletedComicRepository(db)
	ccv := validator.NewCompletedComicValidator()
	completedComicUsecase := usecase.NewCompletedComicUsecase(completedComicRepo, ccv)
	completedComicController := NewCompletedComicController(completedComicUsecase)

	user := testhelper.TestUsers["testuser1"]
	userAuth := testhelper.TestUserAuths["testuser1"]
	if err := userRepo.CreateUser(&user, &userAuth); err != nil {
		t.Fatalf("Error creating user: %v", err)
	}
	if err := completedComicRepo.CreateCompletedComic(&testhelper.TestCompletedComic); err != nil {
		t.Fatalf("Error creating CreateCompletedComic: %v", err)
	}

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/completedComics/user/:userId/:pageNumber")
	c.SetParamNames("userId", "pageNumber")
	c.SetParamValues(user.ID, "1")

	t.Run("ユーザーが読み終えた作品を返す", func(t *testing.T) {
		err := completedComicController.GetUserCompletedComics(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		var completedComicRes []model.CompletedComic
		err = json.Unmarshal(rec.Body.Bytes(), &completedComicRes)
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

	userRepo := repository.NewUserRepository(db)
	followingRepo := repository.NewFollowingRepository(db)
	completedComicRepo := repository.NewCompletedComicRepository(db)
	ccv := validator.NewCompletedComicValidator()
	completedComicUsecase := usecase.NewCompletedComicUsecase(completedComicRepo, ccv)
	completedComicController := NewCompletedComicController(completedComicUsecase)

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
		t.Fatalf("Error creating CreateCompletedComic: %v", err)
	}

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/completedComics/follow/:userId/:pageNumber")
	c.SetParamNames("userId", "pageNumber")
	c.SetParamValues(followUser.ID, "1")

	t.Run("フォローしているユーザーが読み終えた作品を返す", func(t *testing.T) {
		err := completedComicController.GetFolllowCompletedComics(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		var completedComicRes []model.CompletedComic
		err = json.Unmarshal(rec.Body.Bytes(), &completedComicRes)
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

	userRepo := repository.NewUserRepository(db)
	completedComicRepo := repository.NewCompletedComicRepository(db)
	ccv := validator.NewCompletedComicValidator()
	completedComicUsecase := usecase.NewCompletedComicUsecase(completedComicRepo, ccv)
	completedComicController := NewCompletedComicController(completedComicUsecase)

	user := testhelper.TestUsers["testuser1"]
	userAuth := testhelper.TestUserAuths["testuser1"]

	if err := userRepo.CreateUser(&user, &userAuth); err != nil {
		t.Fatalf("Error creating user: %v", err)
	}
	if err := completedComicRepo.CreateCompletedComic(&testhelper.TestCompletedComic); err != nil {
		t.Fatalf("Error creating CreateCompletedComic: %v", err)
	}

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/completedComics/getLatestCompletedComics")

	t.Run("最新の読み終えた作品を返す", func(t *testing.T) {
		err := completedComicController.GetLatestCompletedComics(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		var completedComicRes []model.CompletedComic
		err = json.Unmarshal(rec.Body.Bytes(), &completedComicRes)
		assert.NoError(t, err)
		assert.Equal(t, testhelper.TestCompletedComic.ID, completedComicRes[0].ID)
		assert.Equal(t, testhelper.TestCompletedComic.Rating, completedComicRes[0].Rating)
		assert.Equal(t, testhelper.TestCompletedComic.Review, completedComicRes[0].Review)
		assert.Equal(t, testhelper.TestCompletedComic.UserID, completedComicRes[0].UserID)
		assert.Equal(t, testhelper.TestCompletedComic.ComicID, completedComicRes[0].ComicID)
	})
}
