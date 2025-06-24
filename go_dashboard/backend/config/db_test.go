package config

import (
	"testing"
)

func TestConnectDB(t *testing.T) {
	// Attempt to connect
	ConnectDB()

	// Check if DB is nil
	if DB == nil {
		t.Fatal("DB connection is nil after ConnectDB()")
	}

	// Try pinging the database
	err := DB.Ping()
	if err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}
}