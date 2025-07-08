package dao

import (
	"transactionapp/config"
	"transactionapp/model"
)

func GetAllAccounts() ([]model.Account, error) {
	rows, err := config.DB.Query("SELECT id, name, balance FROM accounts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var accounts []model.Account
	for rows.Next() {
		var a model.Account
		if err := rows.Scan(&a.ID, &a.Name, &a.Balance); err != nil {
			return nil, err
		}
		accounts = append(accounts, a)
	}
	return accounts, nil
}

func GetAccountByID(id int) (model.Account, error) {
	var a model.Account
	err := config.DB.QueryRow("SELECT id, name, balance FROM accounts WHERE id=?", id).
		Scan(&a.ID, &a.Name, &a.Balance)
	return a, err
}
func DeleteAccount(id int) error {
	_, err := config.DB.Exec("DELETE FROM accounts WHERE id=?", id)
	return err
}
func CreateAccount(account model.Account) error {
	_, err := config.DB.Exec("INSERT INTO accounts(name, balance) VALUES(?, ?)", account.Name, account.Balance)
	return err
}
func UpdateAccount(id int, account model.Account) error {
	_, err := config.DB.Exec("UPDATE accounts SET name=?, email=? WHERE id=?", account.Name, account.Balance, id)
	return err
}
