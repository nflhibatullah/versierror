package user

import (
	"github.com/labstack/echo/v4"
	"layered/configs"
	"layered/delivery/middlewares"
	"layered/entities"
	"layered/repository/user"
	"net/http"
)

type UserController struct {
	Repo user.User
}

func New(user user.User) *UserController {
	return &UserController{Repo: user}
}

func (uc UserController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := uc.Repo.GetAll()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Something wrong")
		}
		return c.JSON(
			http.StatusOK, map[string]interface{}{
				"message": "success get all data",
				"data":    users,
			},
		)
	}
}

//func (uc UserController) Get() echo.HandlerFunc {
//	return func(c echo.Context) error {
//		id, _ := strconv.Atoi(c.Param("id"))
//		res, err := uc.Repo.Get(id)
//
//		if err != nil {
//			return c.JSON(http.StatusBadRequest, "User tidak ditemukan atau ada kesalahan pada sistem")
//		}
//
//		return c.JSON(
//			http.StatusOK, map[string]interface{}{
//				"messages": "Berhasil Get data user by id:",
//				"user":     res,
//			},
//		)
//	}
//}

//func (uc UserController) Create() echo.HandlerFunc {
//	return func(c echo.Context) error {
//		tmp := entities.User{}
//		c.Bind(&tmp)
//		res, err := uc.Repo.Create(tmp)
//
//		if err != nil {
//			return c.JSON(http.StatusBadRequest, "Ada kesalahan dalam input")
//		}
//
//		return c.JSON(http.StatusOK, res)
//	}
//}
//
func (uc UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		tmpUser := entities.User{}

		if err := c.Bind(&tmpUser); err != nil {
			return err
		}

		res, err := uc.Repo.Login(tmpUser.Email, tmpUser.Password)
		if err != nil {
			return c.JSON(http.StatusNotFound, "Data tidak ditemukan")
		}

		tmpUser.Token, _ = middlewares.CreateToken(int(tmpUser.ID), configs.SecretKey)
		return c.JSON(
			http.StatusOK, map[string]interface{}{
				"message": "Berhasil login",
				"data":    res,
			},
		)
	}
}

//func (uc UserController) Delete() echo.HandlerFunc {
//	return func(e echo.Context) error {
//		var id, _ = strconv.Atoi(e.Param("id"))
//		res, err := uc.Repo.Delete(id)
//		if err != nil {
//			return e.JSON(http.StatusInternalServerError, "Ada kesalahan")
//		}
//		return e.JSON(
//			http.StatusOK, map[string]interface{}{
//				"messages": "Berhasil menghapus user by id",
//				"id":       res,
//			},
//		)
//	}
//}
//
//func (uc UserController) Update() echo.HandlerFunc {
//	return func(c echo.Context) error {
//		id, _ := strconv.Atoi(c.Param("id"))
//		findUser := entities.User{}
//
//		if err := c.Bind(&findUser); err != nil {
//			return err
//		}
//		res, err := uc.Repo.Update(findUser, id)
//		if err != nil {
//			return c.JSON(http.StatusInternalServerError, "ada kesalahan")
//		}
//
//		return c.JSON(
//			http.StatusOK, map[string]interface{}{
//				"data": res,
//			},
//		)
//
//	}
//}
