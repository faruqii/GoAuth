package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string   `json:"name" validate:"required, string"`
	ProductType string   `json:"product_type" validate:"required, string"`
	Price       float64  `json:"price" validate:"required"`
	MerchantID  int64    `json:"merchant_id"`
	Merchant    Merchant `json:"merchant"`
}

type CreateProductParam struct {
	Name        string  `json:"name" validate:"required, string"`
	ProductType string  `json:"product_type" validate:"required, string"`
	Price       float64 `json:"price" validate:"required"`
	MerchantID  int64   `json:"merchant_id"`
}
