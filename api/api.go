package api

import (
	"errors"
	"net/http"
	"os"
	"strconv"
)

type IComicAPI interface {
	GetGenreComics(genreName string, pageNumber uint) (resp *http.Response, err error)
	GetSearchComics(searchText string, pageNumber uint) (resp *http.Response, err error)
}

type comicAPI struct {
	baseURL string
}

func NewAPI() IComicAPI {

	baseURL := "https://app.rakuten.co.jp/services/api/Kobo/EbookSearch/20170426?applicationId=" + os.Getenv("APPLICATION_ID")

	return &comicAPI{baseURL}
}

func (api *comicAPI) GetGenreComics(genreName string, pageNumber uint) (resp *http.Response, err error) {

	var genreCode string
	switch genreName {
	case "home":
		genreCode = "101904"
	case "shonen":
		genreCode = "101904001"
	case "seinen":
		genreCode = "101904002"
	case "shojo":
		genreCode = "101904011"
	case "ladies":
		genreCode = "101904012"
	default:
		return nil, errors.New("無効なジャンルコード")
	}

	url := api.baseURL + "&koboGenreId=" + genreCode + "&sort=sales&hits=12&page=" + strconv.FormatUint(uint64(pageNumber), 10)

	resp, err = http.Get(url)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (api *comicAPI) GetSearchComics(searchText string, pageNumber uint) (resp *http.Response, err error) {

	url := api.baseURL + "&title=" + searchText + "&sort=sales&hits=10&page=" + strconv.FormatUint(uint64(pageNumber), 10)

	resp, err = http.Get(url)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
