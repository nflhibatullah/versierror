package user

import (
	"gorm.io/gorm"
	"layered/entities"
	"strconv"
)

type UserRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Create(u entities.User) (entities.User, error) {

	if err := ur.db.Save(&u).Error; err != nil {
		return u, err
	}

	return u, nil

}

func (ur *UserRepository) Login(email, password string) ([]entities.User, error) {
	foundUser := []entities.User{}
	var err error
	if err = ur.db.Where("Email = ? AND Password = ?", email, password).Find(&foundUser).Error; err != nil {

		return foundUser, err
	}

	return foundUser, nil
}

func (ur *UserRepository) GetAll() ([]entities.User, error) {
	users := []entities.User{}
	if err := ur.db.Find(&users).Error; err != nil {

		return nil, err
	}
	return users, nil
}

func (ur *UserRepository) Get(UserId int) ([]entities.User, error) {

	user := []entities.User{}

	if err := ur.db.Find(&user, UserId).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) Delete(UserId int) (string, error) {
	var user = []entities.User{}
	ur.db.Delete(&user, UserId)
	return strconv.Itoa(UserId), nil
}

func (ur *UserRepository) Update(newUser entities.User, userId int) (entities.User, error) {
	var oldUser = entities.User{}

	if err := ur.db.Find(&oldUser, "?=id", userId).Error; err != nil {
		return oldUser, err
	}

	oldUser.Nama = newUser.Nama
	oldUser.HP = newUser.HP
	oldUser.Books = newUser.Books
	ur.db.Save(&oldUser)

	return oldUser, nil
}
