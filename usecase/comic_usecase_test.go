package usecase

import (
	"io"
	"net/http"
	"portfolio-backend/model"
	"portfolio-backend/testhelper"
	"portfolio-backend/validator"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockComicRepository struct{}

func (mcr *MockComicRepository) GetGenreComics(genreCode string, pageNumber uint) (*http.Response, error) {
	responseJSON := testhelper.TestComicResponseJSON
	response := &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Body:          io.NopCloser(strings.NewReader(responseJSON)),
		ContentLength: int64(len(responseJSON)),
		Header:        make(http.Header),
	}
	return response, nil
}

func (mcr *MockComicRepository) GetSearchComics(searchText string, pageNumber uint) (*http.Response, error) {
	responseJSON := testhelper.TestComicResponseJSON
	response := &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Body:          io.NopCloser(strings.NewReader(responseJSON)),
		ContentLength: int64(len(responseJSON)),
		Header:        make(http.Header),
	}
	return response, nil
}

func (mcr *MockComicRepository) CreateComics(comics *[]model.Comic) error {
	return nil
}

func TestGetGenreComics(t *testing.T) {
	mockRepo := &MockComicRepository{}
	cv := validator.NewComicValidator()
	cu := NewComicUsecase(mockRepo, cv)

	t.Run("正常系", func(t *testing.T) {
		comics, err := cu.GetGenreComics("genre_code", 1)
		assert.NoError(t, err)
		assert.Equal(t, testhelper.TestComic, comics[0])
	})
}

func TestGetSearchComics(t *testing.T) {
	mockRepo := &MockComicRepository{}
	cv := validator.NewComicValidator()
	cu := NewComicUsecase(mockRepo, cv)

	t.Run("正常系", func(t *testing.T) {
		comics, err := cu.GetSearchComics("search_text", 1)
		assert.NoError(t, err)
		assert.Equal(t, testhelper.TestComic, comics[0])
	})
}
