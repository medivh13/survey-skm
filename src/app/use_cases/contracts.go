package usecases

import (
	loginUc "auth-skm/src/app/use_cases/login"
	userUc "auth-skm/src/app/use_cases/user"
)

type AllUseCases struct {
	UserUseCase  userUc.UserUsecaseInterface
	LoginUseCase loginUc.LoginUsecaseInterface
}
