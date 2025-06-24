package routes

import (
	//"net/http"

	"github.com/MalkiPreethimali/go_dashboard/backend/handlers"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Register your endpoint
	r.HandleFunc("/api/country-revenue", handlers.GetCountryRevenueTable).Methods("GET")
	r.HandleFunc("/api/top-products", handlers.GetTopProducts).Methods("GET")
	r.HandleFunc("/api/monthly-sales", handlers.GetMonthlySales).Methods("GET")
	r.HandleFunc("/api/top-regions", handlers.GetTopRegions).Methods("GET")

	return r

}