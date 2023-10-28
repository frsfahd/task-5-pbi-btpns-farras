package app

import (
	"github.com/go-playground/validator/v10"
)

var valid *validator.Validate

func init() {
	valid = validator.New(validator.WithRequiredStructEnabled())
}

func ValidateUserRegister(data UserRegisterInput) error {
	err := valid.Struct(data)

	return err
}

func ValidateUserUpdate(data UserUpdateInput) error {
	err := valid.Struct(data)

	return err
}

func ValidatePhotoCreate(data PhotoCreateInput) error {
	err := valid.Struct(data)

	return err
}

func ValidatePhotoUpdate(data PhotoUpdateInput) error {
	err := valid.Struct(data)

	return err
}
