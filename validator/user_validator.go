package validator

import (
	"errors"
	"portfolio-backend/model"
)

type IUserValidator interface {
	UserValidate(user model.User) error
	UserAuthValidate(userAuth model.UserAuth) error
}

type userValidator struct{}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}

func (uv *userValidator) UserValidate(user model.User) error {

	if user.ID == "" {
		return errors.New("userIDは空ではいけません")
	}

	if len(user.ID) < 4 || len(user.ID) > 16 {
		return errors.New("userIDは4文字から16文字の範囲内である必要があります")
	}

	if user.UserName == "" {
		return errors.New("userNameは空ではいけません")
	}

	if len(user.UserName) > 16 {
		return errors.New("userNameは16文字の範囲内である必要があります")
	}

	return nil
}

func (uv *userValidator) UserAuthValidate(userAuth model.UserAuth) error {

	if userAuth.Password == "" {
		return errors.New("Passwordは空ではいけません")
	}

	if len(userAuth.Password) < 6 || len(userAuth.Password) > 16 {
		return errors.New("Passwordは6文字から16文字の範囲内である必要があります")
	}

	return nil
}
