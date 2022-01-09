package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"layered/configs"
	"layered/delivery/controllers/user"
)

func RegisterPath(e *echo.Echo, uc *user.UserController) {

	e.Pre(middleware.RemoveTrailingSlash())
	eAuth := e.Group("/akses")
	eAuth.Use(middleware.JWT([]byte(configs.SecretKey)))
	//eAuth.GET("/users/:id", uc.Get())
	e.GET("/users", uc.GetAll(), middleware.JWT([]byte(configs.SecretKey)))
	//eAuth.DELETE("users/:id", uc.Delete())
	//eAuth.PUT("users/:id", uc.Update())

	//e.POST("/users", uc.Create())
	e.POST("/users/login", uc.Login())
}
