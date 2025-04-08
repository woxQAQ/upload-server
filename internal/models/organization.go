package models

type Organization struct {
	Model

	Name        *string `gorm:"not null"`
	DisplayName string
	Avatar      string
	Description string
}
