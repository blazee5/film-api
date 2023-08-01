package models

type Film struct {
	Id          int32 `gorm:"primaryKey"`
	Title       string
	Description string
	Genre       string
}
