package google

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/itolog/go-convertapitos/src/pkg/config"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const USER_URL = "https://www.googleapis.com/oauth2/v1/userinfo"

func ConfigGoogle() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     config.GetConfigEnv("GOOGLE_CLIENT_ID"),
		ClientSecret: config.GetConfigEnv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  config.GetConfigEnv("GOOGLE_REDIRECT_URL"),
		Scopes: []string{
			"email",
			"profile",
		},
		Endpoint: google.Endpoint,
	}
	return conf
}

func GetUser(token *oauth2.Token) (GoogleResponse, error) {
	client := ConfigGoogle().Client(context.Background(), token)

	response, err := client.Get(USER_URL)
	defer response.Body.Close()

	if err != nil {
		return GoogleResponse{}, fmt.Errorf("%v", err)
	}

	var user GoogleResponse

	err = json.NewDecoder(response.Body).Decode(&user)
	if err != nil {
		return GoogleResponse{}, fmt.Errorf("%v", err)
	}
	return user, nil
}
