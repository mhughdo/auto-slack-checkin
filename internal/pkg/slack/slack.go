package slack

import (
	"fmt"
	"net/http"

	"github.com/slack-go/slack"
	"github.com/spf13/viper"
)

type customHttpClient struct {
	http.Client
}

func (c customHttpClient) Do(req *http.Request) (*http.Response, error) {
	cookie := viper.GetString("cookie")
	if cookie == "" {
		return nil, fmt.Errorf("cookie is required")
	}

	req.Header.Add("cookie", cookie)

	return c.Client.Do(req)
}

func NewCustomHTTPClient() customHttpClient {
	return customHttpClient{
		Client: http.Client{},
	}
}

func SendMessage() error {
	token := viper.GetString("token")
	channelID := viper.GetString("channel-id")
	message := viper.GetString("message")

	if token == "" {
		return fmt.Errorf("token is required")
	}

	if channelID == "" {
		return fmt.Errorf("channel-id is required")
	}

	if message == "" {
		return fmt.Errorf("message is required")
	}

	api := slack.New(token, slack.OptionHTTPClient(NewCustomHTTPClient()))

	_, _, err := api.PostMessage(channelID, slack.MsgOptionText(message, false))

	if err != nil {
		if err.Error() == "invalid_auth" {
			fmt.Println("Invalid token or cookie")
		}

		return err
	}

	fmt.Printf("Message sent to channel %s\n", channelID)

	return nil
}
