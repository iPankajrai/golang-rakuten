package models

type User struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name"`
	Gender       string `json:"gender" validate:"required,oneof=Male Female Other"`
	Age          int    `json:"age" validate:"required,min=1"`
	Address      string `json:"address" validate:"required"`
	ContactNo    string `gorm:"unique" json:"contact_no" validate:"required,e164"`
	ContactEmail string `json:"contact_email" validate:"required,email"`
	Photo        string `json:"photo" validate:"url"`
	// Photo        string `json:"photo" validate:"required,url"`
}
