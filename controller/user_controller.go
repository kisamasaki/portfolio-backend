package controller

import (
	"fmt"
	"net/http"
	"os"
	"portfolio-backend/model"
	"portfolio-backend/usecase"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	UpdateUser(c echo.Context) error
	GetUser(c echo.Context) error
	GetLatestUsers(c echo.Context) error
	DeleteUser(c echo.Context) error
	CheckCreateUserStatus(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

type UserData = struct {
	Id       string `json:"id"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func (uc *userController) CheckCreateUserStatus(c echo.Context) error {
	userId := c.Param("userId")
	checkCompletedComicStatus, err := uc.uu.CheckCreateUserStatus(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, checkCompletedComicStatus)
}

func (uc *userController) SignUp(c echo.Context) error {

	userData := UserData{}

	if err := c.Bind(&userData); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user := model.User{}
	user.ID = userData.Id
	user.UserName = userData.UserName

	userAuth := model.UserAuth{}
	userAuth.ID = uuid.New().String()
	userAuth.Password = userData.Password

	err := uc.uu.SignUp(user, userAuth)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (uc *userController) LogIn(c echo.Context) error {

	userData := UserData{}
	if err := c.Bind(&userData); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user := model.User{}
	user.ID = userData.Id
	user.UserName = userData.UserName

	userAuth := model.UserAuth{}
	userAuth.Password = userData.Password

	userResponse, err := uc.uu.Login(user, userAuth)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, userResponse)
}

func (uc *userController) UpdateUser(c echo.Context) error {
	userId := c.Param("userId")
	userName := c.FormValue("userName")
	imageFile, err := c.FormFile("image")

	if err == http.ErrMissingFile {
		println("ファイルのみアップロード")
	} else if err != nil {
		return c.String(http.StatusBadRequest, "ファイルのアップロードエラー")
	} else {
		src, err := imageFile.Open()
		if err != nil {
			return c.String(http.StatusInternalServerError, "画像データの読み込みエラー")
		}
		defer src.Close()

		sess, err := session.NewSession(&aws.Config{
			Region:      aws.String(os.Getenv("AWS_REGION")),
			Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
		})

		if err != nil {
			fmt.Println("セッション生成エラー", err)
			return err
		}

		svc := s3.New(sess)

		_, err = svc.PutObject(&s3.PutObjectInput{
			Bucket:      aws.String(os.Getenv("AWS_BUCKETNAME")),
			Key:         aws.String("users/" + userId + "/avatar.png"),
			Body:        src,
			ContentType: aws.String("image/png"),
		})

		if err != nil {
			fmt.Println("アップロードエラー", err)
			return err
		}
	}

	err = uc.uu.UpdateUser(userId, userName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}

func (uc *userController) GetUser(c echo.Context) error {
	userId := c.Param("userId")
	userResponse, err := uc.uu.GetUser(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, userResponse)
}

func (uc *userController) GetLatestUsers(c echo.Context) error {
	users, err := uc.uu.GetLatestUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

func (uc *userController) DeleteUser(c echo.Context) error {
	userId := c.Param("userId")
	err := uc.uu.DeleteUser(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}
