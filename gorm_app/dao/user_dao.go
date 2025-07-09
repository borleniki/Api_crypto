package dao

import (
	"gormapp/config"
	"gormapp/model"
)

// func GetAllUsers() ([]model.User, error) {
// 	var users []model.User
// 	err := config.DB.Find(&users).Error //select * from users
// 	return users, err
// }

//	func GetUserByID(id int) (model.User, error) {
//		var user model.User
//		err := config.DB.First(&user, id).Error //select * from users where id = ?
//		return user, err
//	}
func GetAllUsers() ([]model.User, error) {
	var users []model.User
	err := config.DB.Preload("Profile").Find(&users).Error
	return users, err
}

func GetUserByID(id int) (model.User, error) {
	var user model.User
	err := config.DB.Preload("Profile").First(&user, id).Error
	return user, err
}
func CreateUser(user model.User) error {
	return config.DB.Create(&user).Error //insert into users (name, email) values (?, ?)
}

func UpdateUser(user model.User) error {
	return config.DB.Save(&user).Error //update users set name = ?, email = ? where id = ?
}

func DeleteUser(id int) error {
	return config.DB.Delete(&model.User{}, id).Error //delete from users where id = ?
}
