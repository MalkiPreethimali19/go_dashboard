package models

type TopProduct struct {
	ProductName   string `json:"product_name"`
	TotalQuantity int    `json:"total_quantity"`
	StockQuantity int    `json:"stock_quantity"`
}