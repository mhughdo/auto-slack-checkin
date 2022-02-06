package main

import (
	"fmt"
	"net/http"

	"github.com/slack-go/slack"
)

const (
	channelID = ""
	token = ""
	dCookie = ""
)


type customHttpClient struct {
	http.Client
}

func (c customHttpClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Add("cookie", dCookie)

	return c.Client.Do(req)
}

func NewCustomHTTPClient() customHttpClient {
	return customHttpClient{
		Client: http.Client{},
	}
}

func main() {
	api := slack.New(token, slack.OptionHTTPClient(NewCustomHTTPClient()))
	// If you set debugging, it will log all requests to the console
	// Useful when encountering issues
	// slack.New("YOUR_TOKEN_HERE", slack.OptionDebug(true))
	a, b, c := api.PostMessage(channelID, slack.MsgOptionText("Hello from bot", false))

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
