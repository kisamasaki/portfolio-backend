package usecase

import (
	"os"
	"portfolio-backend/model"
	"portfolio-backend/repository"
)

type IFollowingUsecase interface {
	Follow(userID, followUserID string) error
	Unfollow(userID, unfollowUserID string) error
	GetFollowingUsers(userID string) ([]model.User, error)
	GetFollowerUsers(userID string) ([]model.User, error)
	CheckFollowingStatus(userID, followUserID string) (bool, error)
}

type followingUsecase struct {
	fr repository.IFollowingRepository
}

func NewFollowingUsecase(fr repository.IFollowingRepository) IFollowingUsecase {
	return &followingUsecase{fr}
}

func (fu *followingUsecase) Follow(userID, followUserID string) error {
	return fu.fr.Follow(userID, followUserID)
}

func (fu *followingUsecase) Unfollow(userID, unfollowUserID string) error {
	return fu.fr.Unfollow(userID, unfollowUserID)
}

func (fu *followingUsecase) GetFollowingUsers(userID string) ([]model.User, error) {
	users, err := fu.fr.GetFollowingUsers(userID)

	if err != nil {
		return []model.User{}, err
	}

	for i := range users {
		users[i].ImageURL = "https://" + os.Getenv("AWS_BUCKETNAME") + ".s3." + os.Getenv("AWS_REGION") + ".amazonaws.com/users/" + users[i].ID + "/avatar.png"
	}

	return users, nil
}

func (fu *followingUsecase) GetFollowerUsers(userID string) ([]model.User, error) {
	users, err := fu.fr.GetFollowerUsers(userID)

	if err != nil {
		return []model.User{}, err
	}

	for i := range users {
		users[i].ImageURL = "https://" + os.Getenv("AWS_BUCKETNAME") + ".s3." + os.Getenv("AWS_REGION") + ".amazonaws.com/users/" + users[i].ID + "/avatar.png"
	}

	return users, nil
}

func (fu *followingUsecase) CheckFollowingStatus(userID, followUserID string) (bool, error) {
	return fu.fr.CheckFollowingStatus(userID, followUserID)
}
