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

type Publisher interface {
	PublishLoginAlert(request LoginRequest) error

	PublishRegisterAlert(request RegisterRequest) error

	PublishActivationAlert(request RegisterRequest) error
}
