package validator

import (
	"errors"
	"portfolio-backend/model"
)

type ICompletedComicValidator interface {
	CompletedComicValidate(completedComic model.CompletedComic) error
}

type completedComicValidator struct{}

func NewCompletedComicValidator() ICompletedComicValidator {
	return &completedComicValidator{}
}

func (ccv *completedComicValidator) CompletedComicValidate(completedComic model.CompletedComic) error {

	if completedComic.Rating > 5 {
		return errors.New("Ratingは5以下である必要があります")
	}

	if len(completedComic.Review) > 100 {
		return errors.New("Reviewは100文字以内である必要があります")
	}

	return nil
}
