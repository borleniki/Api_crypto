package model

type User struct {
	ID    int   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`

	//Profile Profile `gorm:"foreignKey:UserID" json:"profile"` //ONE-TO-ONE relationship

	//Orders []Order `gorm:"foreignKey:UserID" json:"orders"` //ONE-TO-MANY relationship
}