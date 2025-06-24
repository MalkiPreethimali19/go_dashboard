package services

import (
	"testing"

	"github.com/MalkiPreethimali/go_dashboard/backend/config"
)

func setupTestDB() {
	config.ConnectDB()
}

//Test GetCountryRevenue()
func TestGetCountryRevenue(t *testing.T) {
	setupTestDB()
	results, err := GetCountryRevenue()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(results) == 0 {
		t.Log("Warning: Country revenue returned empty. Did you insert test data?")
	}
}

// Test GetTop20Products()
func TestGetTop20Products(t *testing.T) {
	setupTestDB()
	products, err := GetTop20Products()
	if err != nil {
		t.Fatalf("Error fetching top products: %v", err)
	}
	if len(products) > 20 {
		t.Errorf("Expected at most 20 products, got %d", len(products))
	}
}

//Test GetMonthlySalesVolume()
func TestGetMonthlySalesVolume(t *testing.T) {
	setupTestDB()
	sales, err := GetMonthlySalesVolume()
	if err != nil {
		t.Fatalf("Error getting monthly sales: %v", err)
	}
	if len(sales) == 0 {
		t.Log("Note: Monthly sales volume is empty — consider inserting sample transactions")
	}
}

//Test GetTopRegionsByRevenue()
func TestGetTopRegionsByRevenue(t *testing.T) {
	setupTestDB()
	regions, err := GetTopRegionsByRevenue()
	if err != nil {
		t.Fatalf("Error fetching top regions: %v", err)
	}
	if len(regions) > 30 {
		t.Errorf("Expected at most 30 regions, got %d", len(regions))
	}
}