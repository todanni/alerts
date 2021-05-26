package alerts

import "github.com/stretchr/testify/mock"

type MockAlerter struct {
	mock.Mock
}

func (a *MockAlerter) PublishLoginAlert(request LoginRequest) error {
	args := a.Called(request)
	return args.Error(0)
}

func (a *MockAlerter) PublishRegisterAlert(request RegisterRequest) error {
	args := a.Called(request)
	return args.Error(0)
}

func (a *MockAlerter) PublishActivationAlert(request RegisterRequest) error {
	args := a.Called(request)
	return args.Error(0)
}
