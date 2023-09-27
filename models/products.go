package models

import "time"

type Product struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Title     string `json:"title"`
	Price     uint   `json:"price"`
	CreatedAt time.Time
}
