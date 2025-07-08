package model

type Transaction struct {
	ID            int     `json:"id"`
	FromAccountID int     `json:"from_acc"`
	ToAccountID   int     `json:"to_acc"`
	Amount        float64 `json:"amount"`
	CreatedAt     string  `json:"created_at"`
}
