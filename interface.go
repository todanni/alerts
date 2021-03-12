package alerts

type LoginRequest struct {
	Email string
}

type RegisterRequest struct {
	FullName string
	Email string
}

type Alerter interface {
	SendLoginAlert(request LoginRequest) error
	SendRegisterAlert(request RegisterRequest) error
	SendActivationAlert(request RegisterRequest) error
}