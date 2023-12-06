package validator

import "errors"

type IComicValidator interface {
	GetGenreComicsValidate(genreCode string, pageNumber uint) error
	GetSearchComicsValidate(searchText string, pageNumber uint) error
}

type comicValidator struct{}

func NewComicValidator() IComicValidator {
	return &comicValidator{}
}

func (cv *comicValidator) GetGenreComicsValidate(genreCode string, pageNumber uint) error {
	if pageNumber < 1 || pageNumber > 3 {
		return errors.New("pageNumberは1から3の範囲内である必要があります")
	}
	return nil
}

func (cv *comicValidator) GetSearchComicsValidate(searchText string, pageNumber uint) error {
	if searchText == "" {
		return errors.New("searchTextは空ではいけません")
	}

	if pageNumber < 1 || pageNumber > 3 {
		return errors.New("pageNumberは1から3の範囲内である必要があります")
	}

	return nil
}
