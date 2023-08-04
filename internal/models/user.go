package models

type User struct {
	Id       int64 `gorm:"primaryKey"`
	Name     string
	Email    string `gorm:"unique"`
	Password string
}
