package usecase

import (
	"encoding/json"
	"fmt"
	"portfolio-backend/model"
	"portfolio-backend/repository"
	"portfolio-backend/validator"
)

type IComicUsecase interface {
	GetGenreComics(genreCode string, pageNumber uint) ([]model.Comic, error)
	GetSearchComics(searchText string, pageNumber uint) ([]model.Comic, error)
}

type comicUsecase struct {
	cr repository.IComicRepository
	cv validator.IComicValidator
}

func NewComicUsecase(cr repository.IComicRepository, cv validator.IComicValidator) IComicUsecase {
	return &comicUsecase{cr, cv}
}

type Response struct {
	Items []struct {
		Item model.Comic `json:"Item"`
	} `json:"Items"`
}

func (cu *comicUsecase) GetGenreComics(genreCode string, pageNumber uint) ([]model.Comic, error) {
	if err := cu.cv.GetGenreComicsValidate(genreCode, pageNumber); err != nil {
		return nil, err
	}

	comics := []model.Comic{}
	comicResponse, err := cu.cr.GetGenreComics(genreCode, pageNumber)

	if err != nil {
		fmt.Println("ジャンルコミックの取得中にエラー", err)
		return nil, err
	}

	var response Response
	if err := json.NewDecoder(comicResponse.Body).Decode(&response); err != nil {
		fmt.Println("デコード中にエラー", err)
		return nil, err
	}

	for _, v := range response.Items {
		t := model.Comic{
			ID:            v.Item.ID,
			Title:         v.Item.Title,
			Author:        v.Item.Author,
			ItemCaption:   v.Item.ItemCaption,
			LargeImageURL: v.Item.LargeImageURL,
			SalesDate:     v.Item.SalesDate,
		}
		comics = append(comics, t)
	}

	if err := cu.cr.CreateComics(&comics); err != nil {
		return nil, err
	}

	return comics, nil

}

func (cu *comicUsecase) GetSearchComics(searchText string, pageNumber uint) ([]model.Comic, error) {
	if err := cu.cv.GetSearchComicsValidate(searchText, pageNumber); err != nil {
		return nil, err
	}

	comics := []model.Comic{}
	comicResponse, _ := cu.cr.GetSearchComics(searchText, pageNumber)

	var response Response
	if err := json.NewDecoder(comicResponse.Body).Decode(&response); err != nil {
		fmt.Println("デコード中にエラー", err)
	}

	for _, v := range response.Items {
		t := model.Comic{
			ID:            v.Item.ID,
			Title:         v.Item.Title,
			Author:        v.Item.Author,
			ItemCaption:   v.Item.ItemCaption,
			LargeImageURL: v.Item.LargeImageURL,
			SalesDate:     v.Item.SalesDate,
		}
		comics = append(comics, t)
	}

	if err := cu.cr.CreateComics(&comics); err != nil {
		return nil, err
	}

	return comics, nil

}
