package usecases

import (
	loginUc "survey-skm/src/app/use_cases/login"
	userUc "survey-skm/src/app/use_cases/user"
)

type AllUseCases struct {
	UserUseCase  userUc.UserUsecaseInterface
	LoginUseCase loginUc.LoginUsecaseInterface
}
