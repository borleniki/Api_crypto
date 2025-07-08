package service

import (
	"transactionapp/dao"
	"transactionapp/model"
)

func GetAccounts() ([]model.Account, error) {
	return dao.GetAllAccounts()
}

func GetAccount(id int) (model.Account, error) {
	return dao.GetAccountByID(id)
}

func Create(account model.Account) error {
	return dao.CreateAccount(account)
}

func Update(id int, account model.Account) error {
	return dao.UpdateAccount(id, account)
}

func Delete(id int) error {
	return dao.DeleteAccount(id)
}