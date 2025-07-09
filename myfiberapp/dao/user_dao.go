

package dao

import (
	"myfiberapp/config"
	"myfiberapp/model"
)

func GetAllUsers() ([]model.User, error) {
	rows, err := config.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []model.User
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func GetUserByID(id int) (model.User, error) {
	var u model.User
	err := config.DB.QueryRow("SELECT id, name, email FROM users WHERE id=?", id).
		Scan(&u.ID, &u.Name, &u.Email)
	return u, err
}
func DeleteUser(id int) error {
	_, err := config.DB.Exec("DELETE FROM users WHERE id=?", id)
	return err
}
func CreateUser(user model.User) error {
	_, err := config.DB.Exec("INSERT INTO users(name, email) VALUES(?, ?)", user.Name, user.Email)
	return err
}
func UpdateUser(id int, user model.User) error {
	_, err := config.DB.Exec("UPDATE users SET name=?, email=? WHERE id=?", user.Name, user.Email, id)
	return err
}