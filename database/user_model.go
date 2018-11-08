package database

type User struct {
	UserID    int    `gorm:"PRIMARY_KEY" json:"user_id"`
	FirstName string `json:"first_name"`
}
