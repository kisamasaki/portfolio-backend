package validator

import (
	"testing"

	"portfolio-backend/model"

	"github.com/stretchr/testify/assert"
)

func TestCompletedComicValid(t *testing.T) {
	ccv := NewCompletedComicValidator()

	t.Run("正常系", func(t *testing.T) {
		completedComic := model.CompletedComic{
			Rating: 4,
			Review: "記録用",
		}
		err := ccv.CompletedComicValidate(completedComic)
		assert.NoError(t, err)
	})
}

func TestCompletedComicInvalidRating(t *testing.T) {
	ccv := NewCompletedComicValidator()

	t.Run("Ratingが5より大きい", func(t *testing.T) {
		completedComic := model.CompletedComic{
			Rating: 6,
			Review: "記録用",
		}
		err := ccv.CompletedComicValidate(completedComic)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Ratingは5以下である必要があります")
	})
}

func TestGetGenreComicsInvalidReview(t *testing.T) {
	ccv := NewCompletedComicValidator()

	t.Run("Reviewが100文字を超える", func(t *testing.T) {
		completedComic := model.CompletedComic{
			Rating: 4,
			Review: "あああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああ",
		}
		err := ccv.CompletedComicValidate(completedComic)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Reviewは100文字以内である必要があります")
	})
}
