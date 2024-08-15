package repositories

import (
	"user-registration/database"
	"user-registration/models"
)

// create user
func CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

// get user
func GetUser(id uint) (*models.User, error) {
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// update user
func UpdateUser(user *models.User) error {
	return database.DB.Save(user).Error
}

// delete user
func DeleteUser(id uint) error {
	return database.DB.Delete(&models.User{}, id).Error
}

// get All user
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
