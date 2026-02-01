package model

type Product struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	CategoryID int    `json:"category_id"`
}

type ProductJoin struct {
	ID           int    `json:"id"`
	Price        int    `json:"price"`
	Stock        int    `json:"stock"`
	ProductName  string `json:"product_name"`
	CategoryName string `json:"category_name"`
}
