package controller

import (
	"net/http"
	"portfolio-backend/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IComicController interface {
	GetGenreComics(c echo.Context) error
	GetSearchComics(c echo.Context) error
}

type comicController struct {
	cu usecase.IComicUsecase
}

func NewComicController(cu usecase.IComicUsecase) IComicController {
	return &comicController{cu}
}

func (cc *comicController) GetGenreComics(c echo.Context) error {
	genreCode := c.QueryParam("genreCode")

	pageNumber, err := strconv.ParseUint(c.QueryParam("pageNumber"), 10, 0)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "無効なページ番号"})
	}

	comics, err := cc.cu.GetGenreComics(genreCode, uint(pageNumber))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, comics)
}

func (cc *comicController) GetSearchComics(c echo.Context) error {
	searchText := c.Param("searchText")

	pageNumber, err := strconv.ParseUint(c.Param("pageNumber"), 10, 0)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "無効なページ番号"})
	}

	comics, err := cc.cu.GetSearchComics(searchText, uint(pageNumber))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, comics)
}
