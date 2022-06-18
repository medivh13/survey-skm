package user_usecases

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : auth-skm
 */

import (
	dto "auth-skm/src/app/dtos/users"
	"auth-skm/src/domain/repositories"
	"context"
	"log"
)

type UserUsecaseInterface interface {
	Register(ctx context.Context, data *dto.UserReqDTO) (*dto.UserRespDTO, error)
}

type userUseCase struct {
	UserRepo repositories.UsersRepository
}

func NewUserUseCase(userRepo repositories.UsersRepository) *userUseCase {
	return &userUseCase{
		UserRepo: userRepo,
	}
}

func (uc *userUseCase) Register(ctx context.Context, data *dto.UserReqDTO) (*dto.UserRespDTO, error) {

	dataRegister, err := uc.UserRepo.Register(ctx, data)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return dto.ToReturnSaveUser(dataRegister), nil
}
