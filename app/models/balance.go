package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

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
	Last         bool
	Comment      string
}

func (balance Balance) DisplayName() string {
	return fmt.Sprintf("%v (%v)", balance.Product, balance.Price)
}
