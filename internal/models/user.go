package models

type User struct {
	Model

	Name   *string `gorm:"not null"`
	Avatar string
	Email  string
	Phone  string
}
