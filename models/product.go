package models

type Product struct {
	Id          int64    `json:"id"`
	Name        string   `json:"name" validate:"required, string"`
	ProductType string   `json:"product_type" validate:"required, string"`
	Price       int      `json:"price" validate:"required"`
	Merchant    Merchant `gorm:"references:Name" json:"merchant"`
}
