package alerts

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type discordAlerter struct {
	client *http.Client
	webhook string
}

func NewDiscordAlerter(client *http.Client, webhook string) Alerter {
	return &discordAlerter{
		client: client,
		webhook: webhook,
	}
}

func (a discordAlerter) SendLoginAlert(request LoginRequest) error {
	//TODO: Retrieve user details to populate message
	message := Message{
		Embed: []*MessageEmbed{{
			Title:       "A user has just logged in",
			Description: request.Email,
			Image: &MessageEmbedImage{
				URL: "https://i.imgur.com/ibL476q.png",
			},
		},
		}}
	return a.sendMessage(message)
}

func (a discordAlerter) SendRegisterAlert(request RegisterRequest) error {
	message := Message{
		Embed: []*MessageEmbed{{
			Title:       "New register request received",
			Description: request.Email,
		},
		}}

	return a.sendMessage(message)
}

func (a discordAlerter) sendMessage (message Message) error {
	postBody, err := json.Marshal(message)
	if err != nil {
		return err
	}

	b := bytes.NewBuffer(postBody)
	req, err := http.NewRequest(http.MethodPost, a.webhook, b)
	req.Header.Add("Content-Type", "application/json")

	res, err := a.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	log.Infof(string(body))
	return nil
}


func (a discordAlerter) SendActivationAlert() {
	panic("implement me")
}

