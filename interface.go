package alerts

type LoginRequest struct {
	Email string
}

type RegisterRequest struct {
	FullName string
	Email    string
}

type Alert struct {
	Type     string            `json:"type"`
	Source   string            `json:"source"`
	Message  string            `json:"message"`
	Metadata map[string]string `json:"metadata"`
}

type Alerter interface {
	SendLoginAlert(request LoginRequest) error

	SendRegisterAlert(request RegisterRequest) error

	SendActivationAlert(request RegisterRequest) error
}
