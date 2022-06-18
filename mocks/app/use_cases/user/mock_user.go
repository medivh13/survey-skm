package mock_cases

import (
	"context"

	dto "auth-skm/src/app/dtos/users"

	"github.com/stretchr/testify/mock"
)

type MockUserUseCase struct {
	mock.Mock
}

func (m *MockUserUseCase) Register(ctx context.Context, data *dto.UserReqDTO) (*dto.UserRespDTO, error) {
	args := m.Called(data)
	var err error
	var dataResp *dto.UserRespDTO

	if n, ok := args.Get(0).(*dto.UserRespDTO); ok {
		dataResp = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return dataResp, err
}
