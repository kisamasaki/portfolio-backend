package usecase

import (
	"os"
	"portfolio-backend/model"
	"portfolio-backend/repository"
	"portfolio-backend/validator"

	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user model.User, userAuth model.UserAuth) error
	Login(user model.User, userAuth model.UserAuth) (model.User, error)
	UpdateUser(userId string, userName string) error
	GetUser(userId string) (model.User, error)
	GetLatestUsers() ([]model.User, error)
	DeleteUser(userId string) error
	CheckCreateUserStatus(userId string) (bool, error)
}

type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	return &userUsecase{ur, uv}
}

func (uu *userUsecase) CheckCreateUserStatus(userId string) (bool, error) {
	checkCreateUserStatus, err := uu.ur.CheckCreateUserStatus(userId)
	if err != nil {
		return false, err
	}
	return checkCreateUserStatus, nil
}

func (uu *userUsecase) SignUp(user model.User, userAuth model.UserAuth) error {
	if err := uu.uv.UserValidate(user); err != nil {
		return err
	}

	if err := uu.uv.UserAuthValidate(userAuth); err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(userAuth.Password), 10)
	if err != nil {
		return err
	}

	newUser := model.User{
		ID:       user.ID,
		UserName: user.UserName,
		ImageURL: "https://" + os.Getenv("AWS_BUCKETNAME") + ".s3." + os.Getenv("AWS_REGION") + ".amazonaws.com/users/" + user.ID + "/avatar.png",
	}

	newUserAuth := model.UserAuth{
		UserID:   user.ID,
		Password: string(hash),
	}

	if err := uu.ur.CreateUser(&newUser, &newUserAuth); err != nil {
		return err
	}
	return nil
}

func (uu *userUsecase) Login(user model.User, userAuth model.UserAuth) (model.User, error) {

	if err := uu.uv.UserAuthValidate(userAuth); err != nil {
		return model.User{}, err
	}

	storedUser := model.User{}
	if err := uu.ur.GetUser(&storedUser, user.ID); err != nil {
		return model.User{}, err
	}

	storedUserAuth := model.UserAuth{}
	if err := uu.ur.GetUserAuth(&storedUserAuth, user.ID); err != nil {
		return model.User{}, err
	}

	err := bcrypt.CompareHashAndPassword([]byte(storedUserAuth.Password), []byte(userAuth.Password))
	if err != nil {
		return model.User{}, err
	}

	userResponse := model.User{
		ID:       storedUser.ID,
		UserName: storedUser.UserName,
		ImageURL: "https://" + os.Getenv("AWS_BUCKETNAME") + ".s3." + os.Getenv("AWS_REGION") + ".amazonaws.com/users/" + user.ID + "/avatar.png",
	}

	return userResponse, nil
}

func (uu *userUsecase) UpdateUser(userId string, userName string) error {

	if err := uu.ur.UpdateUser(userId, userName); err != nil {
		return err
	}

	return nil
}

func (uu *userUsecase) GetUser(userId string) (model.User, error) {
	storedUser := model.User{}
	if err := uu.ur.GetUser(&storedUser, userId); err != nil {
		return model.User{}, err
	}

	userResponse := model.User{
		ID:         storedUser.ID,
		UserName:   storedUser.UserName,
		ImageURL:   "https://" + os.Getenv("AWS_BUCKETNAME") + ".s3." + os.Getenv("AWS_REGION") + ".amazonaws.com/users/" + storedUser.ID + "/avatar.png",
		Followings: storedUser.Followings,
		Followers:  storedUser.Followers,
	}

	return userResponse, nil
}

func (uu *userUsecase) GetLatestUsers() ([]model.User, error) {
	storedUser := []model.User{}
	if err := uu.ur.GetLatestUsers(&storedUser); err != nil {
		return []model.User{}, err
	}
	for i := range storedUser {
		storedUser[i].ImageURL = "https://" + os.Getenv("AWS_BUCKETNAME") + ".s3." + os.Getenv("AWS_REGION") + ".amazonaws.com/users/" + storedUser[i].ID + "/avatar.png"
	}
	return storedUser, nil
}

func (uu *userUsecase) DeleteUser(userId string) error {
	if err := uu.ur.DeleteUser(userId); err != nil {
		return err
	}

	return nil
}
