package main

import(
	"fmt"
	"log"
	"net/http"

	"go-pg-app/internal/db"
	"go-pg-app/internal/user"
    "github.com/joho/godotenv"
	"github.com/gorilla/mux"
)

func main() {
	if err := godotenv.Load(); err != nil {
        log.Println(".env file not found, relying on system env")
    }

	database, err := db.NewPostgres()

	if err != nil {
		log.Fatal(err)
	}

	if err := database.Ping(); err != nil {
		log.Fatal("DB connection failed: ", err)
	}

	fmt.Println("PostgreSQL connected")

	userRepo := user.NewRepository(database)
    userService := user.NewService(userRepo)
    userHandler := user.NewHandler(userService)

    r := mux.NewRouter()

    r.HandleFunc("/api/users", userHandler.Index).Methods("GET")
    r.HandleFunc("/api/users", userHandler.Create).Methods("POST")
//     r.HandleFunc("/users/{id}", userHandler.Update).Methods("PUT")
//     r.HandleFunc("/users/{id}", userHandler.Delete).Methods("DELETE")

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK Health ishlati hozir ishalti test qilayabman"))
	}).Methods(http.MethodGet)


	log.Println("🚀 Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}