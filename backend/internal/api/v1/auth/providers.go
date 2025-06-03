package auth

import (
	"github.com/itolog/go-convertapitos/backend/pkg/environments"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
)

func setupOAuthProviders() {
	googleClientID := environments.GetEnv("GOOGLE_CLIENT_ID")
	googleClientSecret := environments.GetEnv("GOOGLE_CLIENT_SECRET")
	googleClientRedirect := environments.GetEnv("GOOGLE_REDIRECT_URL")

	githubClientID := environments.GetEnv("GITHUB_CLIENT_ID")
	githubClientSecret := environments.GetEnv("GITHUB_CLIENT_SECRET")
	githubClientRedirect := environments.GetEnv("GITHUB_REDIRECT_URL")

	goth.UseProviders(
		google.New(
			googleClientID,
			googleClientSecret,
			googleClientRedirect,
			"email", "profile",
		),
		github.New(
			githubClientID,
			githubClientSecret,
			githubClientRedirect,
			"user:email",
		),
	)
}
