package services

import (
	"log"
	"github.com/MalkiPreethimali/go_dashboard/backend/config"
	"github.com/MalkiPreethimali/go_dashboard/backend/models"
)

func GetCountryRevenue() ([]models.CountryRevenue, error) {
	query := `
		SELECT
        	r.country,
            p.name AS product_name,
            SUM(t.total_price) AS total_revenue,
            COUNT(t.transaction_id) AS transaction_count
        FROM transactions t
        JOIN customers c ON t.customer_id = c.customer_id
        JOIN regions r ON c.region_id = r.region_id
        JOIN products p ON t.product_id = p.product_id
        GROUP BY r.country, p.name
        ORDER BY total_revenue DESC
	`
	rows, err := config.DB.Query(query)
	if err != nil {
		log.Println("DB query error:", err)
		return nil, err
	}
	defer rows.Close()

	var results []models.CountryRevenue
	for rows.Next() {
		var row models.CountryRevenue
		err := rows.Scan(&row.Country, &row.ProductName, &row.TotalRevenue, &row.TransactionCount)
		if err != nil {
			log.Println("Row scan error:", err)
			continue
		}
		results = append(results, row)
	}

	if err = rows.Err(); err != nil {
		log.Println("Rows error:", err)
		return nil, err
	}
	return results, nil

}

func GetTop20Products() ([]models.TopProduct, error) {
	query := `
		SELECT
			p.name AS product_name,
			SUM(t.quantity) AS total_quantity,
			p.stock_quantity
		FROM transactions t
		JOIN products p ON t.product_id = p.product_id
		GROUP BY p.name, p.stock_quantity
		ORDER BY total_quantity DESC
		LIMIT 20;
	`
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.TopProduct
	for rows.Next() {
		var row models.TopProduct
		if err := rows.Scan(&row.ProductName, &row.TotalQuantity, &row.StockQuantity); err != nil {
			return nil, err
		}
		results = append(results, row)
	}
	return results, nil 
}

func GetMonthlySalesVolume() ([]models.MonthlySales, error){
	query := `
		SELECT
			DATE_FORMAT(transaction_date, '%Y-%m') AS month,
			COUNT(*) AS transaction_count
		FROM transactions
		GROUP BY month
		ORDER BY month;
	`
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.MonthlySales
	for rows.Next() {
		var row models.MonthlySales
		if err := rows.Scan(&row.Month, &row.TransactionCount); err != nil {
			return nil, err
		}
		results = append(results, row)
	}
	return results, nil

}

func GetTopRegionsByRevenue() ([]models.RegionRevenue, error) {
	query := `
		SELECT
			r.name AS region_name,
			SUM(t.total_price) AS total_revenue,
			SUM(t.quantity) AS items_sold
		FROM transactions t
		JOIN customers c ON t.customer_id = c.customer_id
		JOIN regions r ON c.region_id = r.region_id
		GROUP BY r.name
		ORDER BY total_revenue DESC
		LIMIT 30;
	`
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.RegionRevenue
	for rows.Next() {
		var row models.RegionRevenue
		if err := rows.Scan(&row.RegionName, &row.TotalRevenue, &row.ItemsSold); err != nil {
			return nil, err
		}
		results = append(results, row)	
	}
	return results, nil
}
