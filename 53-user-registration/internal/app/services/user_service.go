package services

import (
	"errors"
	"user-registration/internal/app/repositories"
	"user-registration/internal/app/validators"
	"user-registration/pkg/models"
)

func CreateUser(user *models.User) error {
	if err := validators.ValidateUser(user); err != nil {
		return err
	}

	if err := repositories.CreateUser(user); err != nil {
		if errors.Is(err, repositories.ErrUserExists) {
			return err
		}
		return err
	}
	return nil
}

func GetUser(id uint) (*models.User, error) {
	return repositories.GetUser(id)
}

func UpdateUser(user *models.User) error {
	if err := validators.ValidateUser(user); err != nil {
		return err
	}
	if err := repositories.UpdateUser(user); err != nil {
		if errors.Is(err, repositories.ErrUserExists) {
			return err
		}
		return err
	}
	return nil
}

func DeleteUser(id uint) error {
	return repositories.DeleteUser(id)
}

func GetAllUsers() ([]models.User, error) {
	return repositories.GetAllUsers()
}
