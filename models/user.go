package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
}

type UserToken struct {
	gorm.Model
	UserID uint   `json:"user_id"`
	Token  string `json:"token"`
}

type RegisterParam struct {
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"`
}

type LoginParam struct {
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"`
}
