package controller

import (
	"net/http"
	"portfolio-backend/model"
	"portfolio-backend/usecase"
	"strconv"

	"github.com/google/uuid"

	"github.com/labstack/echo/v4"
)

type ICompletedComicController interface {
	GetUserCompletedComics(c echo.Context) error
	GetFolllowCompletedComics(c echo.Context) error
	GetLatestCompletedComics(c echo.Context) error
	CreateCompletedComic(c echo.Context) error
	CheckCompletedComicStatus(c echo.Context) error
}

type completedComicController struct {
	ccu usecase.ICompletedComicUsecase
}

func NewCompletedComicController(ccu usecase.ICompletedComicUsecase) ICompletedComicController {
	return &completedComicController{ccu}
}

func (ccc *completedComicController) CheckCompletedComicStatus(c echo.Context) error {
	userId := c.Param("userId")
	comicId := c.Param("comicId")

	checkCompletedComicStatus, err := ccc.ccu.CheckCompletedComicStatus(userId, comicId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, checkCompletedComicStatus)
}

func (ccc *completedComicController) GetUserCompletedComics(c echo.Context) error {
	userId := c.Param("userId")
	pageNumberStr := c.Param("pageNumber")

	pageNumber, err := strconv.Atoi(pageNumberStr)
	if err != nil {
		println("エラー")
	}

	completedComic, err := ccc.ccu.GetUserCompletedComics(userId, pageNumber)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, completedComic)
}

func (ccc *completedComicController) GetFolllowCompletedComics(c echo.Context) error {
	userId := c.Param("userId")
	pageNumberStr := c.Param("pageNumber")

	pageNumber, err := strconv.Atoi(pageNumberStr)
	if err != nil {
		println("エラー")
	}

	completedComic, err := ccc.ccu.GetFolllowCompletedComics(userId, pageNumber)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, completedComic)
}

func (ccc *completedComicController) GetLatestCompletedComics(c echo.Context) error {
	completedComic, err := ccc.ccu.GetLatestCompletedComics()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, completedComic)
}

func (ccc *completedComicController) CreateCompletedComic(c echo.Context) error {
	userId := c.Param("userId")

	var comic struct {
		ItemNumber string `json:"itemNumber"`
		Rating     uint   `json:"rating"`
		Review     string `json:"review"`
	}

	if err := c.Bind(&comic); err != nil {
		return err
	}

	completedComic := model.CompletedComic{}
	completedComic.ID = uuid.New().String()
	completedComic.Rating = comic.Rating
	completedComic.Review = comic.Review
	completedComic.UserID = userId
	completedComic.ComicID = comic.ItemNumber

	err := ccc.ccu.CreateCompletedComic(completedComic)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, nil)
}
