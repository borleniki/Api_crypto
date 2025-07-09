package model

type Order struct {
	ID    int   `gorm:"primaryKey" json:"id"`
	UserID int  `json:"user_id"`
	Item string `json:"item"`
	Price string `json:"price"`
}