package user

import "layered/entities"

type GetUserResponseFormat struct {
	Data    []entities.User `json:"data"`
	Message string          `json:"message"`
}

type LoginResponseFormat struct {
	Data    entities.User `json:"data"`
	Message string        `json:"message"`
}
