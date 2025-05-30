package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

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
	// feed, err := urlToFeed("https://rss.app/feeds/tQnpi5L0pmGigleL.xml")
	// if err != nil {
	// 	log.Fatal("Error", err)
	// }
	// fmt.Println(feed)
	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL is not found in the environment")
	}
	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	db := database.New(conn)

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	go startScrapping(db, 10, time.Minute)
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
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUserByAPIkey))

	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.handlerGetAllFeeds)
	
	v1Router.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))
	v1Router.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollow))
	v1Router.Delete("/feed_follows/{id}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollow))
	
	v1Router.Get("/posts",apiCfg.middlewareAuth(apiCfg.handleGetPostsForUsers))
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
