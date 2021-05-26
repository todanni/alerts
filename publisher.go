package alerts

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

func NewAlertsPublisher(connString string) (Publisher, error) {
	conn, err := amqp.Dial(connString)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &alertsPublisher{
		channel: ch,
	}, nil
}

type alertsPublisher struct {
	channel *amqp.Channel
}

func (a *alertsPublisher) PublishLoginAlert(request LoginRequest) error {
	al := Alert{
		Type:     "login",
		Source:   "account-service",
		Message:  "New login request received",
		Metadata: make(map[string]string),
	}
	al.Metadata["email"] = request.Email

	b, err := json.Marshal(&al)
	if err != nil {
		return err
	}
	return a.publish(b)
}

func (a *alertsPublisher) PublishRegisterAlert(request RegisterRequest) error {
	al := Alert{
		Type:     "register",
		Source:   "account-service",
		Message:  "New register request received",
		Metadata: make(map[string]string),
	}
	al.Metadata["email"] = request.Email
	al.Metadata["full_name"] = request.FullName

	b, err := json.Marshal(&al)
	if err != nil {
		return err
	}
	return a.publish(b)
}

func (a *alertsPublisher) PublishActivationAlert(request RegisterRequest) error {
	al := Alert{
		Type:     "verify",
		Source:   "account-service",
		Message:  "An account has been verified",
		Metadata: make(map[string]string),
	}
	al.Metadata["email"] = request.Email
	b, err := json.Marshal(&al)
	if err != nil {
		return err
	}
	return a.publish(b)
}

func (a *alertsPublisher) publish(alert []byte) error {
	return a.channel.Publish(
		"",
		"alerts",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        alert,
		})
}
