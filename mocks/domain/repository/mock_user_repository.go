package mock_repository

/*
 * Author      : Jody (jody.almaida@gmail)
 * Modifier    :
 * Domain      : auth-skm
 */

import (
	"context"

	dtoLogin "auth-skm/src/app/dtos/login"
	dto "auth-skm/src/app/dtos/users"
	models "auth-skm/src/infra/models"

	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func (o *MockUserRepo) Register(ctx context.Context, data *dto.UserReqDTO) (*models.Users, error) {
	args := o.Called(data)

	var (
		err      error
		respData *models.Users
	)
	if n, ok := args.Get(0).(*models.Users); ok {
		respData = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return respData, err
}

func (o *MockUserRepo) GetUser(ctx context.Context, data *dtoLogin.LoginReqDTO) (*models.Users, error) {
	args := o.Called(data)

	var (
		err      error
		respData *models.Users
	)
	if n, ok := args.Get(0).(*models.Users); ok {
		respData = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return respData, err
}

func (o *MockUserRepo) UpdateToken(ctx context.Context, id int64, token string) error {
	args := o.Called(id, token)

	var err error

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return err
}
