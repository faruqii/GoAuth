package models

import "gorm.io/gorm"

type Merchant struct {
	gorm.Model
	Name     string    `json:"name"`
	Products []Product `json:"products"`
}

type CreateMerchantParam struct {
	Name string `json:"name"`
}
