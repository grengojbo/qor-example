package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type BalanceApi struct {
	ID        uint    `json:"id"`
	ProductID uint    `json:"product"`
	Count     float32 `json:"count"`
	Price     float32 `json:"price"`
	UserID    uint    `json:"user"`
	StoreID   uint    `json:"store"`
	Comment   string  `json:"comment"`
}

type Balance struct {
	gorm.Model
	ProductID    uint
	Product      Product
	Count        float32
	Price        float32
	UserID       uint
	User         User
	StoreID      uint
	Store        Store
	SubscribedAt time.Time
	Last         bool `sql:"default:false"`
	Comment      string
}

func (balance Balance) DisplayName() string {
	return fmt.Sprintf("%v (%v)", balance.Product, balance.Price)
}
