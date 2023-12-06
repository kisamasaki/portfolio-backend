package repository

import (
	"portfolio-backend/model"

	"gorm.io/gorm"
)

type ICompletedComicRepository interface {
	GetUserCompletedComics(completedComic *[]model.CompletedComic, UserId string, pageNumber int) error
	GetFolllowCompletedComics(completedComic *[]model.CompletedComic, UserId string, pageNumber int) error
	GetLatestCompletedComics(completedComic *[]model.CompletedComic) error
	CreateCompletedComic(completedComic *model.CompletedComic) error
	CheckCompletedComicStatus(userId string, comicId string) (bool, error)
}

type completedComicRepository struct {
	db *gorm.DB
}

func NewCompletedComicRepository(db *gorm.DB) ICompletedComicRepository {
	return &completedComicRepository{db}
}

func (ccr *completedComicRepository) GetUserCompletedComics(completedComic *[]model.CompletedComic, userId string, pageNumber int) error {
	offset := (pageNumber - 1) * 10
	if err := ccr.db.Preload("Comic").Where("user_id = ?", userId).Order("created_at desc").Offset(offset).Limit(10).Find(&completedComic).Error; err != nil {
		return err
	}
	return nil
}

func (ccr *completedComicRepository) CheckCompletedComicStatus(userId string, comicId string) (bool, error) {
	var count int64
	if err := ccr.db.Model(&model.CompletedComic{}).Where("user_id = ? AND comic_id = ?", userId, comicId).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (ccr *completedComicRepository) GetFolllowCompletedComics(completedComic *[]model.CompletedComic, userID string, pageNumber int) error {

	var followerIDs []string
	ccr.db.Model(&model.Following{}).Where("following_id = ?", userID).Pluck("follower_id", &followerIDs)

	offset := (pageNumber - 1) * 10

	ccr.db.Preload("User").Preload("Comic").Where("user_id IN ? OR user_id = ?", followerIDs, userID).Order("created_at desc").Offset(offset).Limit(10).Find(&completedComic)
	return nil
}

func (ccr *completedComicRepository) GetLatestCompletedComics(completedComics *[]model.CompletedComic) error {
	if err := ccr.db.Preload("User").Preload("Comic").Order("created_at desc").Limit(4).Find(&completedComics).Error; err != nil {
		return err
	}
	return nil
}

func (ccr *completedComicRepository) CreateCompletedComic(completedComic *model.CompletedComic) error {
	if err := ccr.db.Create(completedComic).Error; err != nil {
		return err
	}
	return nil
}
