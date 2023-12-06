package main

import (
	"portfolio-backend/api"
	"portfolio-backend/controller"
	"portfolio-backend/db"
	"portfolio-backend/repository"
	"portfolio-backend/router"
	"portfolio-backend/usecase"
	"portfolio-backend/validator"
)

func main() {

	db := db.NewDB()
	api := api.NewAPI()

	userValidator := validator.NewUserValidator()
	comicValidator := validator.NewComicValidator()
	completedComicValidator := validator.NewCompletedComicValidator()

	userRepository := repository.NewUserRepository(db)
	comicRepository := repository.NewComicRepository(db, api)
	completedComicRepository := repository.NewCompletedComicRepository(db)
	userFollowRepository := repository.NewFollowingRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	comicUsecase := usecase.NewComicUsecase(comicRepository, comicValidator)
	completedComicUsecase := usecase.NewCompletedComicUsecase(completedComicRepository, completedComicValidator)
	userFollowUsecase := usecase.NewFollowingUsecase(userFollowRepository)

	userController := controller.NewUserController(userUsecase)
	comicController := controller.NewComicController(comicUsecase)
	completedComicController := controller.NewCompletedComicController(completedComicUsecase)
	userFollowController := controller.NewFollowingController(userFollowUsecase)

	e := router.NewRouter(userController, comicController, completedComicController, userFollowController)
	e.Logger.Fatal(e.Start(":80"))
}
