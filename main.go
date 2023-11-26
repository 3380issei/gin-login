package main

import (
	"gin-login/controller"
	"gin-login/db"
	"gin-login/repository"
	"gin-login/router"
	"gin-login/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	r := router.NewRouter(userController)
	r.Run()
}
