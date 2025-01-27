package main

import (
	"future-path/internal/handler/rest"
	"future-path/internal/repository"
	"future-path/internal/service"
	"future-path/pkg/bcrypt"
	"future-path/pkg/config"
	"future-path/pkg/database/mariadb"
	"future-path/pkg/jwt"
	"future-path/pkg/middleware"
	"future-path/pkg/supabase"
	"log"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

func main() {
	config.LoadEnvironment()
	oauthConfig := config.AuthConfig()
	db, err := mariadb.ConnectDatabase()

	if err != nil {
		log.Fatal(err)
	}

	if err := mariadb.Migrate(db); err != nil {
		log.Fatal(err)
	}

	goth.UseProviders(
		google.New(oauthConfig.GoogleClientID, oauthConfig.GoogleClientSecret, oauthConfig.OAuthCallbackURL),
	)

	repo := repository.NewRepository(db)
	bcrypt := bcrypt.Init()
	jwt := jwt.Init()
	supabase, err := supabase.Init()
	if err != nil {
		log.Fatal(err)
	}
	svc := service.NewService(repo, bcrypt, jwt, supabase)
	middleware := middleware.Init(svc, jwt)

	r := rest.NewRest(svc, middleware)
	r.MountEndpoint()
	r.Run()
}
