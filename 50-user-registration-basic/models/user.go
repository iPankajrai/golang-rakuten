// models/user.go
package models

type User struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Gender       string `json:"gender"`
	Age          int    `json:"age"`
	Address      string `json:"address"`
	ContactNo    string `json:"contact_no"`
	ContactEmail string `json:"contact_email"`
	Photo        string `json:"photo"`
}
