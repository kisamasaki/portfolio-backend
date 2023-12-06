package validator

import (
	"testing"

	"portfolio-backend/model"

	"github.com/stretchr/testify/assert"
)

func TestUserValid(t *testing.T) {
	uv := NewUserValidator()

	t.Run("正常系", func(t *testing.T) {

		user := model.User{
			ID:       "testID",
			UserName: "UserName",
		}

		err := uv.UserValidate(user)
		assert.NoError(t, err)
	})
}

func TestUserInvalidEmptyID(t *testing.T) {
	uv := NewUserValidator()

	t.Run("IDが空", func(t *testing.T) {

		user := model.User{
			ID:       "",
			UserName: "UserName",
		}

		err := uv.UserValidate(user)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "userIDは空ではいけません")

	})
}

func TestUserInvalidLessID(t *testing.T) {
	uv := NewUserValidator()

	t.Run("IDが4文字以下", func(t *testing.T) {

		user := model.User{
			ID:       "aaa",
			UserName: "UserName",
		}

		err := uv.UserValidate(user)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "userIDは4文字から16文字の範囲内である必要があります")

	})
}

func TestUserInvalidGreaterID(t *testing.T) {
	uv := NewUserValidator()

	t.Run("IDが16文字以上", func(t *testing.T) {

		user := model.User{
			ID:       "aaaaaaaaaaaaaaaaa",
			UserName: "UserName",
		}

		err := uv.UserValidate(user)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "userIDは4文字から16文字の範囲内である必要があります")

	})
}

func TestUserInvalidEmptyUserName(t *testing.T) {
	uv := NewUserValidator()

	t.Run("UserNameが空", func(t *testing.T) {

		user := model.User{
			ID:       "testID",
			UserName: "",
		}

		err := uv.UserValidate(user)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "userNameは空ではいけません")

	})
}

func TestUserInvalidGreaterUserName(t *testing.T) {
	uv := NewUserValidator()

	t.Run("UserNameが16文字以上", func(t *testing.T) {

		user := model.User{
			ID:       "testID",
			UserName: "UserNameUserNameUserName",
		}

		err := uv.UserValidate(user)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "userNameは16文字の範囲内である必要があります")

	})
}

func TestUserAuthValid(t *testing.T) {
	uv := NewUserValidator()

	t.Run("正常系", func(t *testing.T) {

		userAuth := model.UserAuth{
			Password: "password",
		}

		err := uv.UserAuthValidate(userAuth)
		assert.NoError(t, err)
	})
}

func TestUserAuthInvalidEmptyPassword(t *testing.T) {
	uv := NewUserValidator()

	t.Run("Passwordが空", func(t *testing.T) {

		userAuth := model.UserAuth{
			Password: "",
		}

		err := uv.UserAuthValidate(userAuth)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Passwordは空ではいけません")

	})
}

func TestUserAuthInvalidLessPassword(t *testing.T) {
	uv := NewUserValidator()

	t.Run("Passwordが6文字以下", func(t *testing.T) {

		userAuth := model.UserAuth{
			Password: "pass",
		}

		err := uv.UserAuthValidate(userAuth)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Passwordは6文字から16文字の範囲内である必要があります")

	})
}

func TestUserAuthInvalidGreaterPassword(t *testing.T) {
	uv := NewUserValidator()

	t.Run("Passwordが16文字以上", func(t *testing.T) {

		userAuth := model.UserAuth{
			Password: "passwordpasswordpassword",
		}

		err := uv.UserAuthValidate(userAuth)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Passwordは6文字から16文字の範囲内である必要があります")

	})
}
