package controller

import (
	"errors"
	"net/http"
	"portfolio-backend/usecase"

	"github.com/labstack/echo/v4"
)

type IFollowingController interface {
	Follow(c echo.Context) error
	Unfollow(c echo.Context) error
	GetFollowingUsers(c echo.Context) error
	GetFollowerUsers(c echo.Context) error
	CheckFollowingStatus(c echo.Context) error
}

type followingController struct {
	fu usecase.IFollowingUsecase
}

func NewFollowingController(fu usecase.IFollowingUsecase) IFollowingController {
	return &followingController{fu}
}

func (fc *followingController) Follow(c echo.Context) error {
	userId := c.Param("userId")
	followerId := c.Param("followerId")

	if userId == "" {
		return errors.New("空になっている")
	}

	err := fc.fu.Follow(userId, followerId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "ユーザーをフォローできませんでした"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "ユーザーをフォローしました"})
}

func (fc *followingController) Unfollow(c echo.Context) error {
	userId := c.Param("userId")
	followerId := c.Param("unfollowUserId")

	err := fc.fu.Unfollow(userId, followerId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to unfollow user"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "フォローを解除しました"})
}

func (fc *followingController) GetFollowingUsers(c echo.Context) error {
	userId := c.Param("userId")
	followingUsers, err := fc.fu.GetFollowingUsers(userId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "フォローユーザーの取得に失敗しました"})
	}

	return c.JSON(http.StatusOK, followingUsers)
}

func (fc *followingController) GetFollowerUsers(c echo.Context) error {
	userId := c.Param("userId")
	followerUsers, err := fc.fu.GetFollowerUsers(userId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "フォロワーの取得に失敗しました"})
	}

	return c.JSON(http.StatusOK, followerUsers)
}

func (fc *followingController) CheckFollowingStatus(c echo.Context) error {
	userId := c.Param("sessionUserId")
	targetUserID := c.Param("userId")

	isFollowing, err := fc.fu.CheckFollowingStatus(userId, targetUserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "フォロー情報の取得に失敗しました",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"isFollowing": isFollowing,
	})
}
