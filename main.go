package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"layered/configs"
	"layered/delivery/controllers/user"
	"layered/delivery/routes"
	_userRepo "layered/repository/user"
	"layered/utils"
)

func main() {
	config := configs.GetConfig()

	db := utils.InitDB(config)
	userRepo := _userRepo.New(db)
	userController := user.New(userRepo)
	//bookRepo := _bookRepo.New(db)
	//bookController := user.New(bookRepo)

	e := echo.New()

	routes.RegisterPath(e, userController)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
