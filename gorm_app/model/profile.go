package model

type Profile struct {
	ID    int   `gorm:"primaryKey" json:"id"`
	UserID int  `json:"user_id"`
	Bio string `json:"bio"`
	Phone string `json:"phone"`
}