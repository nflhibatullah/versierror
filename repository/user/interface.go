package user

import "layered/entities"

type User interface {
	//Get(UserId int) ([]entities.User, error)
	GetAll() ([]entities.User, error)
	//Create(u entities.User) (entities.User, error)
	Login(email, password string) ([]entities.User, error)
	//Delete(UserId int) (string, error)
	//Update(newUser entities.User, userId int) (entities.User, error)
}
