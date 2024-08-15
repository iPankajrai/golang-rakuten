package repositories

import (
	"errors"
	"log"
	db "user-registration/internal/db"

	worker "user-registration/internal/worker"
	models "user-registration/pkg/models"
)

var ErrUserExists = errors.New("user with the same contact number already exists")

func CreateUser(user *models.User) error {
	//check if the user exists with contact number
	var existingUser models.User
	if err := db.DB.Where("contact_no = ?", user.ContactNo).First(&existingUser).Error; err == nil {
		return ErrUserExists
	}

	return db.DB.Create(user).Error
}

func GetUser(id uint) (*models.User, error) {
	log.Printf("Fetching user with ID: %d", id)
	response := worker.QueueReadRequest(id)
	if response.Error != nil {
		log.Printf("Error fetching user with ID: %d, error: %v", id, response.Error)
	}

	return response.User, response.Error

	/*
		without making use of any channels
		var user models.User
		err := db.DB.First(&user, id).Error
		if err != nil {
			log.Printf("Error fetching user with ID: %d, error: %v", id, err)
		}
		return &user, err
	*/
}

func UpdateUser(user *models.User) error {
	var existingUser models.User
	if err := db.DB.Where("(contact_no = ?) AND id != ?", user.ContactNo, user.ID).First(&existingUser).Error; err == nil {
		return ErrUserExists
	}

	return db.DB.Save(user).Error
}

func DeleteUser(id uint) error {
	return db.DB.Delete(&models.User{}, id).Error
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := db.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
