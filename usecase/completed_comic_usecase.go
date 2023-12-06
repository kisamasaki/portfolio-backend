package usecase

import (
	"os"
	"portfolio-backend/model"
	"portfolio-backend/repository"
	"portfolio-backend/validator"
)

type ICompletedComicUsecase interface {
	GetUserCompletedComics(userId string, pageNumber int) ([]model.CompletedComic, error)
	GetFolllowCompletedComics(userId string, pageNumber int) ([]model.CompletedComic, error)
	GetLatestCompletedComics() ([]model.CompletedComic, error)
	CreateCompletedComic(readComic model.CompletedComic) error
	CheckCompletedComicStatus(userId string, comicId string) (bool, error)
}

type completedComicUsecase struct {
	ccr repository.ICompletedComicRepository
	ccv validator.ICompletedComicValidator
}

func NewCompletedComicUsecase(ccr repository.ICompletedComicRepository, ccv validator.ICompletedComicValidator) ICompletedComicUsecase {
	return &completedComicUsecase{ccr, ccv}
}

func (ccu *completedComicUsecase) CheckCompletedComicStatus(userId string, comicId string) (bool, error) {
	checkCompletedComicStatus, err := ccu.ccr.CheckCompletedComicStatus(userId, comicId)
	if err != nil {
		return false, err
	}
	return checkCompletedComicStatus, nil
}

func (ccu *completedComicUsecase) GetFolllowCompletedComics(userId string, pageNumber int) ([]model.CompletedComic, error) {
	completedComic := []model.CompletedComic{}
	if err := ccu.ccr.GetFolllowCompletedComics(&completedComic, userId, pageNumber); err != nil {
		return nil, err
	}
	for i := range completedComic {
		user := &completedComic[i].User
		user.ImageURL = "https://" + os.Getenv("AWS_BUCKETNAME") + ".s3." + os.Getenv("AWS_REGION") + ".amazonaws.com/users/" + user.ID + "/avatar.png"
	}
	return completedComic, nil
}

func (ccu *completedComicUsecase) GetLatestCompletedComics() ([]model.CompletedComic, error) {
	completedComic := []model.CompletedComic{}
	if err := ccu.ccr.GetLatestCompletedComics(&completedComic); err != nil {
		return []model.CompletedComic{}, err
	}

	for i := range completedComic {
		user := &completedComic[i].User
		user.ImageURL = "https://" + os.Getenv("AWS_BUCKETNAME") + ".s3." + os.Getenv("AWS_REGION") + ".amazonaws.com/users/" + user.ID + "/avatar.png"
	}
	return completedComic, nil
}

func (ccu *completedComicUsecase) GetUserCompletedComics(userId string, pageNumber int) ([]model.CompletedComic, error) {
	comics := []model.CompletedComic{}
	if err := ccu.ccr.GetUserCompletedComics(&comics, userId, pageNumber); err != nil {
		return nil, err
	}
	return comics, nil
}

func (ccu *completedComicUsecase) CreateCompletedComic(completedComic model.CompletedComic) error {
	if err := ccu.ccv.CompletedComicValidate(completedComic); err != nil {
		return err
	}

	if err := ccu.ccr.CreateCompletedComic(&completedComic); err != nil {
		return err
	}
	return nil
}
