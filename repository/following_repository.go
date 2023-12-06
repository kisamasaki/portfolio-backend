package repository

import (
	"portfolio-backend/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IFollowingRepository interface {
	Follow(userID, followUserID string) error
	Unfollow(userID, unfollowUserID string) error
	GetFollowingUsers(userID string) ([]model.User, error)
	GetFollowerUsers(userID string) ([]model.User, error)
	CheckFollowingStatus(userID, followUserID string) (bool, error)
}

type followingRepository struct {
	db *gorm.DB
}

func NewFollowingRepository(db *gorm.DB) IFollowingRepository {
	return &followingRepository{db}
}

func (fr *followingRepository) Follow(userID, followerUserID string) error {
	follow := model.Following{
		ID:          uuid.New().String(),
		FollowingID: userID,
		FollowerID:  followerUserID,
	}
	return fr.db.Create(&follow).Error
}

func (fr *followingRepository) Unfollow(userID, unfollowUserID string) error {
	return fr.db.Where("following_id = ? AND follower_id = ?", userID, unfollowUserID).Delete(&model.Following{}).Error
}

func (fr *followingRepository) GetFollowingUsers(userID string) ([]model.User, error) {
	var user []model.User

	if err := fr.
		db.
		Preload("Followings").
		Preload("Followers").
		Joins("INNER JOIN followings ON users.id = followings.follower_id").
		Where("followings.following_id = ?", userID).
		Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (fr *followingRepository) GetFollowerUsers(userID string) ([]model.User, error) {
	var user []model.User

	if err := fr.
		db.
		Preload("Followings").
		Preload("Followers").
		Joins("INNER JOIN followings ON users.id = followings.following_id").
		Where("followings.follower_id = ?", userID).
		Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (fr *followingRepository) CheckFollowingStatus(userID, followUserID string) (bool, error) {
	var count int64
	err := fr.db.
		Model(&model.Following{}).
		Where("following_id = ? AND follower_id = ?", userID, followUserID).
		Count(&count).
		Error
	return count > 0, err
}
