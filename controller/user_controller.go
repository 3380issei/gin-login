package controller

import (
	"gin-login/entity"
	"gin-login/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Signup(c *gin.Context)
	Login(c *gin.Context)
	// LogOut(c *gin.Context) error
}

type userController struct {
	uu usecase.UserUsecase
}

func NewUserController(uu usecase.UserUsecase) UserController {
	return &userController{uu}
}

func (uc *userController) Signup(c *gin.Context) {
	user := entity.User{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	userRes, err := uc.uu.Signup(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, userRes)
}

func (uc *userController) Login(c *gin.Context) {
	user := entity.User{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	tokenString, err := uc.uu.Login(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	cookie := &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		MaxAge:   10,
		Path:     "/",
		Domain:   "localhost",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}

	c.SetCookie(
		cookie.Name,
		cookie.Value,
		cookie.MaxAge,
		cookie.Path,
		cookie.Domain,
		cookie.Secure,
		cookie.HttpOnly,
	)

	c.JSON(http.StatusOK, tokenString)
}

// func (uc *userController) LogOut(c *gin.Context) error {
// 	cookie := new(http.Cookie)
// 	cookie.Name = "token"
// 	cookie.Value = ""
// 	cookie.Expires = time.Now()
// 	cookie.Path = "/"
// 	cookie.Domain = os.Getenv("API_DOMAIN")
// 	cookie.Secure = true
// 	cookie.HttpOnly = true
// 	cookie.SameSite = http.SameSiteNoneMode
// 	c.SetCookie(cookie)
// 	return c.NoContent(http.StatusOK)
// }
