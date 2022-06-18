package users_dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type UsersInterface interface {
	Validate() error
}

type UserReqDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	PassWord string `json:"password"`
}

func (dto *UserReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Email, validation.Required, is.Email),
		validation.Field(&dto.PassWord, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type UserRespDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
