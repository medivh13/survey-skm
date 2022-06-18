package users_handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	mockDTOlogin "auth-skm/mocks/app/dtos/login"
	mockDTOuser "auth-skm/mocks/app/dtos/users"
	mockLoginUC "auth-skm/mocks/app/use_cases/login"
	mockUserUC "auth-skm/mocks/app/use_cases/user"
	mockResp "auth-skm/mocks/interface/rest/response"
	dtoLogin "auth-skm/src/app/dtos/login"
	dtoUser "auth-skm/src/app/dtos/users"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockCheckoutHandler struct {
	mock.Mock
}

type HandlerTest struct {
	suite.Suite
	mockUserUseCase  *mockUserUC.MockUserUseCase
	mockLoginUsecase *mockLoginUC.MockLoginUseCase
	mockResponse     *mockResp.MockResponse
	h                UserHandlerInterface
	w                *httptest.ResponseRecorder
	dtoTest          *dtoUser.UserReqDTO
	dtoTestFail      *dtoUser.UserReqDTO
	mockDTO          *mockDTOuser.MockUserDTO
	dtoLoginTest     *dtoLogin.LoginReqDTO
	mockDTOlogin     *mockDTOlogin.MockLoginDTO
	dtoLoginrsp      *dtoLogin.LoginRespDTO
}

func (suite *HandlerTest) SetupTest() {
	suite.mockUserUseCase = new(mockUserUC.MockUserUseCase)
	suite.mockResponse = new(mockResp.MockResponse)
	suite.mockDTO = new(mockDTOuser.MockUserDTO)
	suite.mockDTOlogin = new(mockDTOlogin.MockLoginDTO)
	suite.h = NewUserHandler(suite.mockResponse, suite.mockUserUseCase, suite.mockLoginUsecase)
	suite.w = httptest.NewRecorder()
	suite.dtoTest = &dtoUser.UserReqDTO{
		Name:     "Test",
		Email:    "email@email.com",
		PassWord: "Password",
	}
	suite.dtoTestFail = nil
	suite.dtoLoginTest = &dtoLogin.LoginReqDTO{
		Email:    "email@email.com",
		PassWord: "test",
	}
	suite.dtoLoginrsp = &dtoLogin.LoginRespDTO{
		Token: "testToken",
	}
}

func (s *HandlerTest) TestRegisterSuccess() {
	bodyBytes, _ := json.Marshal(s.dtoTest)

	r := httptest.NewRequest("POST", "/register/", bytes.NewBuffer(bodyBytes))
	s.mockDTO.Mock.On("Validate").Return(nil)
	s.mockUserUseCase.Mock.On("Register", s.dtoTest).Return(mock.Anything, nil)
	s.mockResponse.Mock.On("JSON", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	http.HandlerFunc(s.h.Register).ServeHTTP(s.w, r)

	s.Equal(200, s.w.Result().StatusCode)

}

func (s *HandlerTest) TestRegisterFailDTO() {
	r := httptest.NewRequest("POST", "/register/", nil)
	s.mockDTO.Mock.On("Validate").Return(errors.New(mock.Anything))
	s.mockResponse.Mock.On("HttpError", mock.Anything, mock.Anything).Return(mock.Anything)

	http.HandlerFunc(s.h.Register).ServeHTTP(s.w, r)

	s.Equal(400, s.w.Result().StatusCode)

}

func (s *HandlerTest) TestRegisterFailValidateDTO() {
	bodyBytes, _ := json.Marshal(s.dtoTestFail)
	r := httptest.NewRequest("POST", "/register/", bytes.NewBuffer(bodyBytes))
	s.mockDTO.Mock.On("Validate").Return(errors.New(mock.Anything))
	s.mockResponse.Mock.On("HttpError", mock.Anything, mock.Anything).Return(mock.Anything)

	http.HandlerFunc(s.h.Register).ServeHTTP(s.w, r)

	s.Equal(400, s.w.Result().StatusCode)

}

func (s *HandlerTest) TestRegisterFail() {
	bodyBytes, _ := json.Marshal(s.dtoTest)

	r := httptest.NewRequest("POST", "/register/", bytes.NewBuffer(bodyBytes))
	s.mockDTO.Mock.On("Validate").Return(nil)
	s.mockUserUseCase.Mock.On("Register", s.dtoTest).Return(nil, errors.New(mock.Anything))
	s.mockResponse.Mock.On("HttpError", mock.Anything, mock.Anything).Return(mock.Anything)

	http.HandlerFunc(s.h.Register).ServeHTTP(s.w, r)

	s.Equal(500, s.w.Result().StatusCode)

}

func (s *HandlerTest) TestLoginFailDTO() {
	// bodyBytes, _ := json.Marshal(s.dtoLoginTest)

	r := httptest.NewRequest("POST", "/signin/", nil)
	s.mockDTO.Mock.On("Validate").Return(errors.New(mock.Anything))
	s.mockResponse.Mock.On("HttpError", mock.Anything, mock.Anything).Return(mock.Anything)
	s.mockResponse.Mock.On("JSON", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	http.HandlerFunc(s.h.Login).ServeHTTP(s.w, r)

	s.Equal(400, s.w.Result().StatusCode)

}

func (s *HandlerTest) TestLoginFailValidateDTO() {
	bodyBytes, _ := json.Marshal(s.dtoTestFail)
	r := httptest.NewRequest("POST", "/signin/", bytes.NewBuffer(bodyBytes))
	s.mockDTO.Mock.On("Validate").Return(errors.New(mock.Anything))
	s.mockResponse.Mock.On("HttpError", mock.Anything, mock.Anything).Return(mock.Anything)
	s.mockResponse.Mock.On("JSON", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	http.HandlerFunc(s.h.Login).ServeHTTP(s.w, r)

	s.Equal(400, s.w.Result().StatusCode)

}

func (s *HandlerTest) TestLoginFail() {
	bodyBytes, _ := json.Marshal(s.dtoLoginTest)

	r := httptest.NewRequest("POST", "/signin/", bytes.NewBuffer(bodyBytes))
	s.mockDTO.Mock.On("Validate").Return(nil)
	s.mockLoginUsecase.Mock.On("Login", s.dtoLoginTest).Return(nil, errors.New(mock.Anything))
	s.mockResponse.Mock.On("HttpError", mock.Anything, mock.Anything).Return(mock.Anything)

	http.HandlerFunc(s.h.Login).ServeHTTP(s.w, r)

	s.Equal(401, s.w.Result().StatusCode)

}

func TestHandler(t *testing.T) {
	suite.Run(t, new(HandlerTest))
}
