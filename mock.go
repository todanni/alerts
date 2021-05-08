package alerts

import "github.com/stretchr/testify/mock"

type AlerterMock struct {
	mock.Mock
}

func (a *AlerterMock) SendLoginAlert(request LoginRequest) error {
	args := a.Called(request)
	return args.Error(0)
}

func (a *AlerterMock) SendRegisterAlert(request RegisterRequest) error {
	args := a.Called(request)
	return args.Error(0)
}

func (a *AlerterMock) SendActivationAlert(request RegisterRequest) error {
	args := a.Called(request)
	return args.Error(0)
}
