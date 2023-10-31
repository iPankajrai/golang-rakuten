package models

type Contact struct { // contacts in the table name it is automatically created in db by GORM
	ID           uint   `json:"id" gorm:"primaryKey"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
	Status       string `json:"status"`
	LastModified int64  `json:"lastModified" gorm:"column:last_modified"`
}
