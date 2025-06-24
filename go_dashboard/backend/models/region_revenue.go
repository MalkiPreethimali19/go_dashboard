package models

type RegionRevenue struct {
	RegionName   string  `json:"region_name"`
	TotalRevenue float64 `json:"total_revenue"`
	ItemsSold    int     `json:"items_sold"`
}