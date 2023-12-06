package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGenreComicsValid(t *testing.T) {
	cv := NewComicValidator()

	t.Run("正常系", func(t *testing.T) {
		err := cv.GetGenreComicsValidate("101904", 2)
		assert.NoError(t, err)
	})
}

func TestGetGenreComicsInvalidLessPageNumber(t *testing.T) {
	cv := NewComicValidator()

	t.Run("ページナンバーが1未満", func(t *testing.T) {
		err := cv.GetGenreComicsValidate("101904", 0)
		assert.Error(t, err)
		assert.Equal(t, "pageNumberは1から3の範囲内である必要があります", err.Error())
	})
}

func TestGetGenreComicsInvalidGreaterPageNumber(t *testing.T) {
	cv := NewComicValidator()

	t.Run("ページナンバーが3より大きい", func(t *testing.T) {
		err := cv.GetGenreComicsValidate("101904", 4)
		assert.Error(t, err)
		assert.Equal(t, "pageNumberは1から3の範囲内である必要があります", err.Error())
	})
}

func TestGetSearchComicsValid(t *testing.T) {
	cv := NewComicValidator()

	t.Run("正常系", func(t *testing.T) {
		err := cv.GetSearchComicsValidate("someSearchText", 2)
		assert.NoError(t, err)
	})
}

func TestGetSearchComicsInvalidEmptySearchText(t *testing.T) {
	cv := NewComicValidator()

	t.Run("searchTextが空", func(t *testing.T) {
		err := cv.GetSearchComicsValidate("", 2)
		assert.Error(t, err)
		assert.Equal(t, "searchTextは空ではいけません", err.Error())
	})
}

func TestGetSearchComicsInvalidLessPageNumber(t *testing.T) {
	cv := NewComicValidator()

	t.Run("ページナンバーが1未満", func(t *testing.T) {
		err := cv.GetSearchComicsValidate("someSearchText", 0)
		assert.Error(t, err)
		assert.Equal(t, "pageNumberは1から3の範囲内である必要があります", err.Error())
	})
}

func TestGetSearchComicsInvalidGreaterPageNumber(t *testing.T) {
	cv := NewComicValidator()

	t.Run("ページナンバーが3より大きい", func(t *testing.T) {
		err := cv.GetSearchComicsValidate("someSearchText", 4)
		assert.Error(t, err)
		assert.Equal(t, "pageNumberは1から3の範囲内である必要があります", err.Error())
	})
}
