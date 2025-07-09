package service

import (
	"gormapp/dao"
	"gormapp/model"
)

func GetUsers() ([]model.User, error) {
	return dao.GetAllUsers()
}

func GetUser(id int) (model.User, error) {
	return dao.GetUserByID(id)
}

func Create(user model.User) error {
	return dao.CreateUser(user)
}

func Update(user model.User) error {
	return dao.UpdateUser(user)
}

func Delete(id int) error {
	return dao.DeleteUser(id)
}
