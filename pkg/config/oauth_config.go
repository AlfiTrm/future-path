package config

import "os"

type OAuthConfig struct {
	GoogleClientID     string
	GoogleClientSecret string
	OAuthCallbackURL   string
}

func AuthConfig() *OAuthConfig {
	return &OAuthConfig{
		GoogleClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		GoogleClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		OAuthCallbackURL:   os.Getenv("REDIRECT_URL"),
	}
}
