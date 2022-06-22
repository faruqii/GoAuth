package models


type Product struct {
	Id int64 `json:"id"` 
	Name string `json:"name" validate:"required, Name"`
	Product_Type string `json:"product_type" validate:"required, Product"`
	Price float64 `json:"price" validate:"required, Price"`
	// Merchant Merchant
}