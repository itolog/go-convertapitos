package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/itolog/go-convertapitos/src/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

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

func GetUser(token string) GoogleResponse {
	reqURL, err := url.Parse("https://www.googleapis.com/oauth2/v1/userinfo")
	if err != nil {
		fmt.Println(err)
	}
	ptoken := fmt.Sprintf("Bearer %s", token)
	res := &http.Request{
		Method: "GET",
		URL:    reqURL,
		Header: map[string][]string{
			"Authorization": {ptoken},
		},
	}
	req, err := http.DefaultClient.Do(res)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	var data GoogleResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	return data
}
