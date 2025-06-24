package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MalkiPreethimali/go_dashboard/backend/config"
)

// Run once to ensure DB is connected
func setupTestEnv() {
	config.ConnectDB()
}

func TestRouterRoutes(t *testing.T) {
	setupTestEnv()

	router := SetupRouter()

	tests := []struct {
		name       string
		method     string
		path       string
		wantStatus int
	}{
		{"Country Revenue", "GET", "/api/country-revenue", http.StatusOK},
		{"Top Products", "GET", "/api/top-products", http.StatusOK},
		{"Monthly Sales", "GET", "/api/monthly-sales", http.StatusOK},
		{"Top Regions", "GET", "/api/top-regions", http.StatusOK},
		{"Not Found", "GET", "/api/does-not-exist", http.StatusNotFound},
	}

	for _, tt := range tests {
		req, err := http.NewRequest(tt.method, tt.path, nil)
		if err != nil {
			t.Errorf("Test %s: Failed to create request: %v", tt.name, err)
			continue
		}
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if rr.Code != tt.wantStatus {
			t.Errorf("Test %s: Expected status %d, got %d", tt.name, tt.wantStatus, rr.Code)
		}
	}
}