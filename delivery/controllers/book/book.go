package book

//
//import (
//	"github.com/labstack/echo/v4"
//	"layered/delivery/controllers/user"
//	"layered/entities"
//	"layered/repository/book"
//	"net/http"
//	"strconv"
//)
//
//type BookController struct {
//	Repo book.Book
//}
//
//func New(Book book.Book) *BookController {
//	return &BookController{Repo: Book}
//}
//
//func (bc BookController) GetAll() echo.HandlerFunc {
//	return func(c echo.Context) error {
//		books, err := bc.Repo.GetAll()
//
//		if err != nil {
//			return c.JSON(http.StatusInternalServerError, "Something wrong")
//		}
//		return c.JSON(
//			http.StatusOK, map[string]interface{}{
//				"message": "success get all data",
//				"data":    books,
//			},
//		)
//	}
//}
//func (uc user.UserController) Get() echo.HandlerFunc {
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
//func (uc user.UserController) Create() echo.HandlerFunc {
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
//func (uc user.UserController) Delete() echo.HandlerFunc {
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
//func (uc user.UserController) Update() echo.HandlerFunc {
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
