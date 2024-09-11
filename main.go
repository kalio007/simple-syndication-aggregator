package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("no port specified uisng defalut port 8000")
		portString = "8080"
	}
	fmt.Printf("Port: %v\n", portString)

	router := chi.NewRouter()
	//adding cors
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	//mounting routes
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)

	//mounting routes
	router.Mount("/v1", v1Router)

	srv := &http.Server{Handler: router, Addr: ":" + portString}

	log.Printf("Server Running on Port : %v\n", portString)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
