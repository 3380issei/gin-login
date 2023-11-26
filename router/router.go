package router

import (
	"gin-login/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter(uc controller.UserController) *gin.Engine {
	r := gin.Default()

	r.POST("/signup", uc.Signup)
	r.POST("/login", uc.Login)
	// r.POST("logout", uc.Logout)

	return r
}
