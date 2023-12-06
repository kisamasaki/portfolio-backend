package repository

import (
	"fmt"
	"os"
	"portfolio-backend/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(user *model.User, userAuth *model.UserAuth) error
	GetUser(user *model.User, userId string) error
	GetUserAuth(userAuth *model.UserAuth, userId string) error
	GetLatestUsers(user *[]model.User) error
	UpdateUser(userID string, userName string) error
	DeleteUser(userId string) error
	CheckCreateUserStatus(userId string) (bool, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUser(user *model.User, userId string) error {

	if err := ur.db.
		Where("id = ?", userId).
		Preload("Followings", "following_id = ?", userId).
		Preload("Followers", "follower_id = ?", userId).
		First(&user).Error; err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) CheckCreateUserStatus(userId string) (bool, error) {
	var count int64

	if err := ur.db.Model(&model.User{}).Where("id = ?", userId).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (ur *userRepository) CreateUser(user *model.User, userAuth *model.UserAuth) error {

	if err := ur.db.Create(&user).Error; err != nil {
		return err
	}

	if err := ur.db.Create(&userAuth).Error; err != nil {
		return err
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
	})
	if err != nil {
		fmt.Println("セッション生成エラー", err)
		return err
	}

	svc := s3.New(sess)

	copyInput := &s3.CopyObjectInput{
		Bucket:     aws.String(os.Getenv("AWS_BUCKETNAME")),
		CopySource: aws.String(os.Getenv("AWS_BUCKETNAME") + "/common/placehold.png"),
		Key:        aws.String("users/" + user.ID + "/avatar.png"),
	}

	_, copyErr := svc.CopyObject(copyInput)
	if copyErr != nil {
		fmt.Println("コピーエラー", copyErr)
		return err
	}

	return nil
}

func (ur *userRepository) UpdateUser(userID string, userName string) error {
	var user model.User
	if err := ur.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return err
	}

	if userName != "" {
		user.UserName = userName
	}

	if err := ur.db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) GetUserAuth(userAuth *model.UserAuth, userId string) error {
	if err := ur.db.Where("user_id=?", userId).First(&userAuth).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) GetLatestUsers(users *[]model.User) error {

	if err := ur.db.Order("created_at desc").Limit(8).Find(&users).Error; err != nil {
		return err
	}

	for i := range *users {
		user := &(*users)[i]
		if err := ur.db.Model(user).Preload("Followings").Preload("Followers").Where("id = ?", user.ID).Find(&user).Error; err != nil {
			return err
		}
	}

	return nil
}

func (ur *userRepository) DeleteUser(userId string) error {
	if err := ur.db.Where("user_id = ?", userId).Delete(&model.CompletedComic{}).Error; err != nil {
		return err
	}

	if err := ur.db.Where("follower_id = ? OR following_id = ?", userId, userId).Delete(&model.Following{}).Error; err != nil {
		return err
	}

	if err := ur.db.Delete(&model.UserAuth{}, "user_id = ?", userId).Error; err != nil {
		return err
	}

	if err := ur.db.Where("id = ?", userId).Delete(&model.User{}).Error; err != nil {
		return err
	}

	return nil
}
