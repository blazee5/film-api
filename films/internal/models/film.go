package models

type Film struct {
	Id          int64 `gorm:"primaryKey"`
	Title       string
	Description string
	Genre       string
}
