package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/MalkiPreethimali/go_dashboard/backend/services"
)

func GetCountryRevenueTable(w http.ResponseWriter, r *http.Request){
	results, err := services.GetCountryRevenue()
	if err != nil {
        http.Error(w, "Error retrieving country revenue data", http.StatusInternalServerError)
        return
    }

	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(results)
}

func GetTopProducts(w http.ResponseWriter, r *http.Request) {
    results, err := services.GetTop20Products()
    if err != nil {
		http.Error(w, "Error fetching top products", http.StatusInternalServerError)
		return
	}

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(results)
}

func GetMonthlySales(w http.ResponseWriter, r *http.Request) {
    results, err := services.GetMonthlySalesVolume()
    if err != nil {
        http.Error(w, "Error fetching monthly sales volume", http.StatusInternalServerError)
        return    
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(results)
}

func GetTopRegions(w http.ResponseWriter, r *http.Request) {
    results, err := services.GetTopRegionsByRevenue()
    if err != nil {
		http.Error(w, "Error fetching top regions", http.StatusInternalServerError)
		return
	}
    
    w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}