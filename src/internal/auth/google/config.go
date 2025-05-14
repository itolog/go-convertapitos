package google

import (
	"github.com/itolog/go-convertapitos/src/pkg/config"
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
