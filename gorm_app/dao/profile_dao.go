package dao

import (
	"gormapp/config"
	"gormapp/model"
)

func GetAllProfiles() ([]model.Profile, error) {
	var profiles []model.Profile
	err := config.DB.Find(&profiles).Error
	return profiles, err
}

func GetProfileByID(id int) (model.Profile, error) {
	var profile model.Profile
	err := config.DB.First(&profile, id).Error
	return profile, err
}
func CreateProfile(profile model.Profile) error {
	return config.DB.Create(&profile).Error //insert into profiles (name, email) values (?, ?)
}

func UpdateProfile(profile model.Profile) error {
	return config.DB.Save(&profile).Error //update profiles set name = ?, email = ? where id = ?
}

func DeleteProfile(id int) error {
	return config.DB.Delete(&model.Profile{}, id).Error //delete from profiles where id = ?
}
