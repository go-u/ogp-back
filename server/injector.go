package main

import (
	"github.com/go-chi/cors"
	"log"
	"os"
	"server/application/usecase"
	"server/infrastructure/auth"
	"server/infrastructure/store"
	"server/infrastructure/twitter"
	"server/interfaces/api"
)

func getProjectId() string {
	PROJECT_ID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if PROJECT_ID == "" {
		log.Fatalln("Failed to Get PROJECT_ID 'GOOGLE_CLOUD_PROJECT'\n If this is local test, Set 'appname-local' as GOOGLE_CLOUD_PROJECT")
	}
	log.Println("PROJECT_ID: ", PROJECT_ID)
	return PROJECT_ID
}

func getApiConfig() api.Config {
	projectID := getProjectId()

	// infra
	sqlHandler := store.NewSqlHandler(projectID)
	authCient := auth.NewClient(projectID)
	twitterClient := twitter.NewClient()

	// repository & service
	authService := auth.NewAuthService(*authCient)
	twitterService := twitter.NewTwitterService(*twitterClient)
	bookmarkRepository := store.NewBookmarkRepository(*sqlHandler)
	ogpRepository := store.NewOgpRepository(*sqlHandler)
	statRepository := store.NewStatRepository(*sqlHandler)
	userRepository := store.NewUserRepository(*sqlHandler)

	// usecase
	authUsecase := usecase.NewAuthUsecase(authService)
	bookmarkUsecase := usecase.NewBookmarkUsecase(bookmarkRepository)
	ogpUsecase := usecase.NewOgpUsecase(ogpRepository, twitterService)
	statUsecase := usecase.NewStatUsecase(statRepository)
	userUsecase := usecase.NewUserUsecase(userRepository)

	config := api.Config{
		AuthUsecase:     authUsecase,
		BookmarkUsecase: bookmarkUsecase,
		OgpUsecase:      ogpUsecase,
		StatUsecase:     statUsecase,
		UserUsecase:     userUsecase,
	}
	return config
}

func getCorsOptions() cors.Options {
	return cors.Options{
		AllowedOrigins: []string{"*"}, // for development
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Content-Type", "JWT"},
		MaxAge:         3600, // Maximum value not ignored by any of major browsers
	}
}
