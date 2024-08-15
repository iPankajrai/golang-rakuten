package validators

import (
	"user-registration/pkg/models"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateUser(user *models.User) error {
	return validate.Struct(user)
}
