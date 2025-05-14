package configs

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func ConfigGoogle() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     GetConfigEnv("GOOGLE_CLIENT_ID"),
		ClientSecret: GetConfigEnv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  GetConfigEnv("GOOGLE_REDIRECT_URL"),
		Scopes: []string{
			"email",
			"profile",
		},
		Endpoint: google.Endpoint,
	}
	return conf
}
