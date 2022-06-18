package mock_logindto

import (
	dto "auth-skm/src/app/dtos/login"

	"github.com/stretchr/testify/mock"
)

type MockLoginDTO struct {
	mock.Mock
}

func NewMockItemDTO() *MockLoginDTO {
	return &MockLoginDTO{}
}

var _ dto.LoginInterface = &MockLoginDTO{}

func (m *MockLoginDTO) Validate() error {
	args := m.Called()
	var err error
	if n, ok := args.Get(0).(error); ok {
		err = n
		return err
	}

	return nil
}
