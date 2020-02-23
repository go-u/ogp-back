package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/volatiletech/sqlboiler/boil"
	"log"
	"net/http"
	"os"
	"server/api/v1"
	"server/initialize"
	"server/initialize/db"
	"server/initialize/firebaseauth"
	"server/initialize/gcs"
	"server/middleware_custom"
)

func init() {
	PROJECT_ID := initialize.GetProjectId()
	// db
	db.InitializeDB(PROJECT_ID)
	// GCS
	gcs.InitializeGcs(PROJECT_ID)
	// firebase
	firebaseauth.InitializeFirebaseAuth(PROJECT_ID)
}

func main() {
	r := chi.NewRouter()
	boil.DebugMode = false
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Content-Type", "JWT"},
		MaxAge:         3600, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)
	r.Use(middleware.Logger)
	r.Use(middleware_custom.LimitSizeByContentLengthHeader)
	r.Use(middleware_custom.GetCloudFlareIp)
	r.Use(middleware.Recoverer)
	r.Mount("/api/v1", v1.ApiRouter())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	log.Printf("Listening on port %s", port)

	http.ListenAndServe(":"+port, r)
	defer db.Connection.Close()
}
