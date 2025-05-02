package main

import (
	"log"
	"net/http"
	"todo-app/internal/handlers"
	"todo-app/internal/middleware"
	"todo-app/pkg/auth"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Initialize router
	r := mux.NewRouter()

	// Initialize JWT service
	jwtService := auth.NewJWTService()

	// CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	// Public routes
	r.HandleFunc("/login", handlers.Login(jwtService)).Methods("POST", "OPTIONS")

	// Protected routes
	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.AuthMiddleware(jwtService))

	// TODO List routes
	protected.HandleFunc("/todos", handlers.GetTodoLists).Methods("GET")
	protected.HandleFunc("/todos", handlers.CreateTodoList).Methods("POST")
	protected.HandleFunc("/todos/{id}", handlers.GetTodoList).Methods("GET")
	protected.HandleFunc("/todos/{id}", handlers.UpdateTodoList).Methods("PUT")
	protected.HandleFunc("/todos/{id}", handlers.DeleteTodoList).Methods("DELETE")

	// TODO Item routes
	protected.HandleFunc("/todos/{listId}/items", handlers.GetTodoItems).Methods("GET")
	protected.HandleFunc("/todos/{listId}/items", handlers.CreateTodoItem).Methods("POST")
	protected.HandleFunc("/todos/{listId}/items/{itemId}", handlers.UpdateTodoItem).Methods("PUT")
	protected.HandleFunc("/todos/{listId}/items/{itemId}", handlers.DeleteTodoItem).Methods("DELETE")

	// Start server with CORS
	handler := c.Handler(r)
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
