package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/volatiletech/sqlboiler/boil"
	"log"
	"net/http"
	"os"
	"server/interfaces/api"
	middlewareCustom "server/interfaces/api/middleware"
)

func init() {
	boil.DebugMode = false
}

func main() {
	// router
	r := chi.NewRouter()
	r.Use(cors.New(getCorsOptions()).Handler)
	r.Use(middleware.Logger)
	r.Use(middlewareCustom.LimitSizeByContentLengthHeader)
	r.Use(middlewareCustom.GetCloudFlareIp)
	r.Use(middleware.Recoverer)

	// Dependency Injection
	handlerConfig := getApiConfig()
	apiHandler := api.NewHandler(&handlerConfig)
	r.Mount("/api", apiHandler)

	// start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	log.Printf("Listening on port %s", port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatalln(err)
	}
}
