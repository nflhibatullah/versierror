package user

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/magiconair/properties/assert"
	"layered/entities"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {
	t.Run(
		"Login user", func(t *testing.T) {
			e := echo.New()
			requestBody, _ := json.Marshal(
				map[string]string{
					"email":    "naufal@0i['trhnaufal.com",
					"password": "123",
				},
			)

			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
			res := httptest.NewRecorder()
			req.Header.Set("Content-Type", "application/json")
			context := e.NewContext(req, res)
			context.SetPath("/users/login")
			UController := New(mockUserRepository{})
			err := UController.Login()(context)
			if err != nil {
				log.Info(err)
			}

			var Responses LoginResponseFormat
			//bagian ini kak yang bikin pass terus
			err = json.Unmarshal(res.Body.Bytes(), &Responses)
			if err != nil {
				return
			}

			assert.Equal(t, Responses.Data.Email, "naufal@naufal.com")
		},
	)
}

type mockUserRepository struct{}

func (m mockUserRepository) GetAll() ([]entities.User, error) {
	return []entities.User{
		{Nama: "naufal", HP: "08123456789"},
	}, nil
}

func (m mockUserRepository) Login(email, password string) ([]entities.User, error) {

	return []entities.User{
		{Email: "naufal@naufal.com", Password: "123"},
	}, nil
}
