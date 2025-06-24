package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MalkiPreethimali/go_dashboard/backend/config"
)

func setupTestDB() {
	config.ConnectDB()
}

func TestGetCountryRevenueTable(t *testing.T) {
	// Ensure DB is connected
	setupTestDB()

	// Create a mock request
	req, err := http.NewRequest("GET", "/api/country-revenue", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Call the handler
	handler := http.HandlerFunc(GetCountryRevenueTable)
	handler.ServeHTTP(rr, req)

	// Assert status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", status)
	}

	// Optional: Check JSON body is not empty
	if rr.Body.Len() == 0 {
		t.Errorf("Expected response body, got empty")
	}
}

func TestGetTopProducts(t *testing.T) {
	setupTestDB()

	req, err := http.NewRequest("GET", "/api/top-products", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(GetTopProducts)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", rr.Code)
	}
	if rr.Body.Len() == 0 {
		t.Error("Expected response body, got empty")
	}
}

func TestGetMonthlySales(t *testing.T) {
	setupTestDB()

	req, err := http.NewRequest("GET", "/api/monthly-sales", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(GetMonthlySales)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", rr.Code)
	}
	if rr.Body.Len() == 0 {
		t.Error("Expected response body, got empty")
	}
}

func TestGetTopRegions(t *testing.T) {
	setupTestDB()

	req, err := http.NewRequest("GET", "/api/top-regions", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(GetTopRegions)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", rr.Code)
	}
	if rr.Body.Len() == 0 {
		t.Error("Expected response body, got empty")
	}
}
