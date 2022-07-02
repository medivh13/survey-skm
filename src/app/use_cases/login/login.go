package user_usecases

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : survey-skm
 */

import (
	"context"
	"log"
	dto "survey-skm/src/app/dtos/login"
	"survey-skm/src/domain/repositories"
	infraHelper "survey-skm/src/infra/helpers"
)

type LoginUsecaseInterface interface {
	Login(ctx context.Context, data *dto.LoginReqDTO) (*dto.LoginRespDTO, error)
}

type loginUseCase struct {
	UserRepo repositories.UsersRepository
}

func NewUserUseCase(userRepo repositories.UsersRepository) *loginUseCase {
	return &loginUseCase{
		UserRepo: userRepo,
	}
}

func (uc *loginUseCase) Login(ctx context.Context, data *dto.LoginReqDTO) (*dto.LoginRespDTO, error) {
	var dataResp dto.LoginRespDTO
	dataUser, err := uc.UserRepo.GetUser(ctx, data)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err = infraHelper.ComparePassword(dataUser.Password, data.PassWord); err != nil {
		return nil, err
	}

	token, err := infraHelper.GenerateJWT()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = uc.UserRepo.UpdateToken(ctx, dataUser.ID, token)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	dataResp.Token = token

	return &dataResp, nil
}
