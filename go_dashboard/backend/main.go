package main

import (
	"log"
	"net/http"

	"github.com/MalkiPreethimali/go_dashboard/backend/config"
	"github.com/MalkiPreethimali/go_dashboard/backend/routes"
)

func main() {
	config.ConnectDB()
	router := routes.SetupRouter() 

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", enableCORS(router)))
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}