package models

type MonthlySales struct {
	Month             string `json:"month"` // format: "2025-06"
	TransactionCount  int    `json:"transaction_count"`
}