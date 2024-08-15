package services

import (
	"user-registration/models"
	"user-registration/repositories"
)

func CreateUser(user *models.User) error {
	return repositories.CreateUser(user)
}

func GetUser(id uint) (*models.User, error) {
	return repositories.GetUser(id)
}

func UpdateUser(user *models.User) error {
	return repositories.UpdateUser(user)
}

func DeleteUser(id uint) error {
	return repositories.DeleteUser(id)
}

func GetAllUsers() ([]models.User, error) {
	return repositories.GetAllUsers()
}
