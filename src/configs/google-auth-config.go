package configs

import (
	"github.com/itolog/go-convertapitos/src/pkg/environments"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func ConfigGoogle() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     environments.GetEnv("GOOGLE_CLIENT_ID"),
		ClientSecret: environments.GetEnv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  environments.GetEnv("GOOGLE_REDIRECT_URL"),
		Scopes: []string{
			"email",
			"profile",
		},
		Endpoint: google.Endpoint,
	}
	return conf
}
