package user_usecases

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : survey-skm
 */

import (
	"context"
	"errors"
	"testing"
	"time"

	mockDTOuser "survey-skm/mocks/app/dtos/users"
	mockUserRepo "survey-skm/mocks/domain/repository"

	dtoUser "survey-skm/src/app/dtos/users"
	models "survey-skm/src/infra/models"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockUsecase struct {
	mock.Mock
}

type UsecaseUserTest struct {
	suite.Suite
	userRepo *mockUserRepo.MockUserRepo

	usecase     UserUsecaseInterface
	userModels  *models.Users
	dtoTest     *dtoUser.UserReqDTO
	dtoTestFail *dtoUser.UserReqDTO
	mockDTO     *mockDTOuser.MockUserDTO
}

func (suite *UsecaseUserTest) SetupTest() {
	suite.userRepo = new(mockUserRepo.MockUserRepo)
	suite.mockDTO = new(mockDTOuser.MockUserDTO)
	suite.usecase = NewUserUseCase(suite.userRepo)

	suite.userModels = &models.Users{
		ID:        1,
		Name:      "Test",
		Email:     "Email",
		Password:  "Password",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	suite.dtoTest = &dtoUser.UserReqDTO{
		Name:     "Test",
		Email:    "Email",
		PassWord: "Password",
	}

}

func (uc *UsecaseUserTest) TestRegisterSuccess() {
	uc.userRepo.Mock.On("Register", uc.dtoTest).Return(uc.userModels, nil)
	_, err := uc.usecase.Register(context.Background(), uc.dtoTest)
	uc.Equal(nil, err)
}

func (uc *UsecaseUserTest) TestCheckoutFailRetrieveItems() {
	uc.userRepo.Mock.On("Register", uc.dtoTest).Return(nil, errors.New(mock.Anything))

	_, err := uc.usecase.Register(context.Background(), uc.dtoTest)
	uc.Error(errors.New(mock.Anything), err)
}

func TestUsecase(t *testing.T) {
	suite.Run(t, new(UsecaseUserTest))
}
