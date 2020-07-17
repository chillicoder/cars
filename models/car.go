// models/car.go

package models

import (
	_ "github.com/jinzhu/gorm"
)

type Car struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Description string `json:"description"`
	Location    string `json:"location"`
}

type CreateCarInput struct {
	Description string `json:"description" binding:"required"`
	Location    string `json:"location" binding:"required"`
}

type UpdateCarInput struct {
	Description string `json:"description"`
	Location    string `json:"location"`
}
