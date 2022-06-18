package mock_userdto

import (
	dto "auth-skm/src/app/dtos/users"

	"github.com/stretchr/testify/mock"
)

type MockUserDTO struct {
	mock.Mock
}

func NewMockItemDTO() *MockUserDTO {
	return &MockUserDTO{}
}

var _ dto.UsersInterface = &MockUserDTO{}

func (m *MockUserDTO) Validate() error {
	args := m.Called()
	var err error
	if n, ok := args.Get(0).(error); ok {
		err = n
		return err
	}

	return nil
}
