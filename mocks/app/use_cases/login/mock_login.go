package mock_cases

import (
	"context"

	dto "survey-skm/src/app/dtos/login"

	"github.com/stretchr/testify/mock"
)

type MockLoginUseCase struct {
	mock.Mock
}

func (m *MockLoginUseCase) Login(ctx context.Context, data *dto.LoginReqDTO) (*dto.LoginRespDTO, error) {
	args := m.Called(data)
	var err error
	var dataResp *dto.LoginRespDTO

	if n, ok := args.Get(0).(*dto.LoginRespDTO); ok {
		dataResp = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return dataResp, err
}
