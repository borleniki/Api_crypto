package service

import (
	"gormapp/dao"
	"gormapp/model"
)

func GetProfiles() ([]model.Profile, error) {
	return dao.GetAllProfiles()
}

func GetProfile(id int) (model.Profile, error) {
	return dao.GetProfileByID(id)
}

func CreateProfile(user model.Profile) error {
	return dao.CreateProfile(user)
}

func UpdateProfile(user model.Profile) error {
	return dao.UpdateProfile(user)
}

func DeleteProfile(id int) error {
	return dao.DeleteProfile(id)
}
