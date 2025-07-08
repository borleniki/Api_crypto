package dao

import (
	"errors"
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
	_, err := config.DB.Exec("UPDATE accounts SET name=?, balance=? WHERE id=?", account.Name, account.Balance, id)
	return err
}
func TransferAmount(fromID, toID int, amount float64) error {
	tx, err := config.DB.Begin()
	if err != nil {
		return err
	}

	var fromBalance float64
	err = tx.QueryRow("SELECT balance FROM accounts WHERE id = ?", fromID).Scan(&fromBalance)
	if err != nil {
		tx.Rollback()
		return err
	}

	if fromBalance < amount {
		tx.Rollback()
		return errors.New("insufficient funds")
	}

	_, err = tx.Exec("UPDATE accounts SET balance = balance - ? WHERE id = ?", amount, fromID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("UPDATE accounts SET balance = balance + ? WHERE id = ?", amount, toID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("INSERT INTO transactions (from_acc, to_acc, amount) VALUES (?, ?, ?)", fromID, toID, amount)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func MiniStatement(accountID int) ([]model.Transaction, error) {
	rows, err := config.DB.Query(`
        SELECT id, from_acc, to_acc, amount, created_at
        FROM transactions
        WHERE from_acc = ? OR to_acc = ?
        ORDER BY created_at DESC LIMIT 5`, accountID, accountID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var txns []model.Transaction
	for rows.Next() {
		var txn model.Transaction
		if err := rows.Scan(&txn.ID, &txn.FromAccountID, &txn.ToAccountID, &txn.Amount, &txn.CreatedAt); err != nil {
			return nil, err
		}
		txns = append(txns, txn)
	}

	return txns, nil
}
