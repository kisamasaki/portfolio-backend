package repository

import (
	"net/http"
	"portfolio-backend/api"
	"portfolio-backend/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IComicRepository interface {
	GetGenreComics(genreCode string, pageNumber uint) (*http.Response, error)
	GetSearchComics(searchText string, pageNumber uint) (*http.Response, error)
	CreateComics(comics *[]model.Comic) error
}

type comicRepository struct {
	db       *gorm.DB
	comicAPI api.IComicAPI
}

func NewComicRepository(db *gorm.DB, comicAPI api.IComicAPI) IComicRepository {
	return &comicRepository{db, comicAPI}
}

func (cr *comicRepository) GetGenreComics(genreCode string, pageNumber uint) (*http.Response, error) {

	comics, err := cr.comicAPI.GetGenreComics(genreCode, pageNumber)

	if err != nil {
		return nil, err
	}

	return comics, nil
}

func (cr *comicRepository) GetSearchComics(searchText string, pageNumber uint) (*http.Response, error) {

	comics, err := cr.comicAPI.GetSearchComics(searchText, pageNumber)

	if err != nil {
		return nil, err
	}

	return comics, nil
}

func (cr *comicRepository) CreateComics(comics *[]model.Comic) error {

	tx := cr.db.Begin()

	if err := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoNothing: true,
	}).Create(comics).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
