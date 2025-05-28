package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Rushikesh-purohit-0503/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL is not found in the environment")
	}
	conn, err:= sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	
	apiCfg := apiConfig{
		DB: database.New(conn),
	}
	router := chi.NewRouter()
	router.Use(
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"*"}, // Allow all origins
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"*"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300,
		}),
	)

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handleErrors)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Addr:    ":" + portString,
		Handler: router,
	}

	log.Printf("Server running on port %v", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server:", err)
	}

}
