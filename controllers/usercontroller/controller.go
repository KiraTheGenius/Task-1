package usercontroller

import (
	"fmt"
	"net/http"
	"ticket/models"
	"ticket/services/userService"
	"ticket/utils"

	"github.com/labstack/echo"
)

type UserController struct {
	UserService userService.UserService
}

type loginReq struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (u *UserController) Signup(c echo.Context) error {
	signupReq := &loginReq{}
	c.Bind(&signupReq)
	// todo check is user name is already registered
	newUser := models.User{
		Username: signupReq.UserName,
		Password: signupReq.Password,
	}
	err := u.UserService.CreateUser(&newUser)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusCreated, nil)
}

func (u *UserController) Login(c echo.Context) error {
	loginReq := &loginReq{}
	c.Bind(&loginReq)
	user, err := u.UserService.GetUser(loginReq.UserName)
	if err != nil {
		return echo.ErrInternalServerError
	}
	fmt.Println(utils.HashPassword(loginReq.Password))
	if utils.ValidatePassword(user.Password, loginReq.Password) {
		token, err := utils.GenerateTokenPair(user)
		if err != nil {
			return echo.ErrInternalServerError
		}
		return c.JSON(http.StatusOK, token)
	}
	return echo.ErrUnauthorized
}
