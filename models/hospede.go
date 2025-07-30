package models

import "gorm.io/gorm"


type Hospede struct{
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name" binding:"required"`
	Birthday string `json:"birthday" binding:"required"`
	City string `json:"city"`
	State string `json:"state"`
	Country string `json:"country"`
	Phone string `json:"phone"`
}

func Migate(db *gorm.DB){
	db.AutoMigrate(&Hospede{})
}
