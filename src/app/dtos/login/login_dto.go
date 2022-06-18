package login_dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type LoginInterface interface {
	Validate() error
}

type LoginReqDTO struct {
	Email    string `json:"email"`
	PassWord string `json:"password"`
}

func (dto *LoginReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Email, validation.Required, is.Email),
		validation.Field(&dto.PassWord, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type LoginRespDTO struct {
	Token string `json:"token"`
}
