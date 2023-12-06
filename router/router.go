package router

import (
	"os"
	"portfolio-backend/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NoCache(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Response().Header().Set("Pragma", "no-cache")
		c.Response().Header().Set("Expires", "0")
		return next(c)
	}
}

func NewRouter(uc controller.IUserController, cc controller.IComicController, rc controller.ICompletedComicController, fc controller.IFollowingController) *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},

		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},

		AllowMethods: []string{"GET", "PUT", "POST", "DELETE"},

		AllowCredentials: true,
	}))

	e.Use(NoCache)

	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.GET("/checkCreateUserStatus/:userId", uc.CheckCreateUserStatus)
	e.GET("/:userId", uc.GetUser)
	e.PUT("/:userId", uc.UpdateUser)
	e.DELETE("/:userId", uc.DeleteUser)
	e.GET("/latestUsers", uc.GetLatestUsers)

	c := e.Group("/comics")
	c.GET("", cc.GetGenreComics)
	c.GET("/:searchText/:pageNumber", cc.GetSearchComics)

	o := e.Group("/completedComics")
	o.GET("/follow/:userId/:pageNumber", rc.GetFolllowCompletedComics)
	o.POST("/:userId", rc.CreateCompletedComic)
	o.GET("/user/:userId/:pageNumber", rc.GetUserCompletedComics)
	o.GET("/getLatestCompletedComics", rc.GetLatestCompletedComics)
	o.GET("/check-completedcomic/:comicId/:userId", rc.CheckCompletedComicStatus)

	f := e.Group("/follow")
	f.POST("/:userId/:followerId", fc.Follow)
	f.POST("/unfollow/:userId/:unfollowUserId", fc.Unfollow)
	f.GET("/following/:userId", fc.GetFollowingUsers)
	f.GET("/followers/:userId", fc.GetFollowerUsers)
	f.GET("/check-follow/:sessionUserId/:userId", fc.CheckFollowingStatus)

	return e
}
