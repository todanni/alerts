package alerts

import "github.com/stretchr/testify/mock"

type MockAlerter struct {
	mock.Mock
}

func (a *MockAlerter) SendLoginAlert(request LoginRequest) error {
	args := a.Called(request)
	return args.Error(0)
}

func (a *MockAlerter) SendRegisterAlert(request RegisterRequest) error {
	args := a.Called(request)
	return args.Error(0)
}

func (a *MockAlerter) SendActivationAlert(request RegisterRequest) error {
	args := a.Called(request)
	return args.Error(0)
}
