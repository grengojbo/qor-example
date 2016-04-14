package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/qor/media_library"
	"github.com/qor/transition"
)

type InvoiceIn struct {
	gorm.Model
	Name           string `json:"name"`
	OrganizationID uint
	Organization   Organization
	StoreID        uint
	Store          Store
	Amount         float32 `json:"summa"`
	ShippedAt      *time.Time
	CancelledAt    *time.Time
	Document       media_library.FileSystem
	Invoices       []InvoiceInItem `json:"invoices"`
	transition.Transition
}

func (self *InvoiceIn) BeforeSave(tx *gorm.DB) (err error) {
	// amount := float32(0)
	// self.Amount = self.Quantity * self.Price
	for _, item := range self.Invoices {
		if item.Price == float32(0) || item.Price < float32(0) {
			item.Price = float32(1)
		}
		item.Amount = item.Quantity * item.Price
		self.Amount += item.Amount
	}
	// fmt.Println("-------------> Invoice in Amount:", self.Amount)
	return nil
}

type InvoiceInItem struct {
	gorm.Model
	InvoiceInID uint
	Code        string `json:"code"`
	ProductID   uint
	Product     Product
	Name        string  `json:"name"`
	Quantity    float32 `json:"count"`
	Price       float32 `json:"price"`
	Amount      float32 `json:"summa"`
}

func (self *InvoiceInItem) BeforeSave(tx *gorm.DB) (err error) {
	if self.Amount == float32(0) {
		self.Amount = self.Quantity * self.Price
		fmt.Println("-------------> InvoiceInItem Amount:", self.Amount)
	}
	return nil
}

type InvoiceOut struct {
	gorm.Model
	Name           string  `json:"name"`
	Partner        string  `json:"partner"`
	Amount         float32 `json:"summa"`
	OrganizationID uint
	Organization   Organization
	StoreID        uint
	Store          Store
	ShippedAt      *time.Time
	CancelledAt    *time.Time
	Document       media_library.FileSystem
	Invoices       []InvoiceOutItem `json:"invoices"`
	transition.Transition
}

type InvoiceOutItem struct {
	gorm.Model
	InvoiceOutID uint
	Code         string `json:"code"`
	ProductID    uint
	Product      Product
	Name         string  `json:"name"`
	Quantity     float32 `json:"count"`
	Price        float32 `json:"price"`
	Amount       float32 `json:"summa"`
}

type Invoice struct {
	ID         uint          `json:"id"`
	Name       string        `json:"name"`
	AutoNumber string        `json:"autoNumber"`
	Partner    string        `json:"partner"`
	Cards      []string      `json:"cards"`
	Barcode    bool          `json:"barcode"`
	Amount     float32       `json:"summa"`
	Invoices   []InvoiceList `json:"invoices"`
}

type InvoiceList struct {
	Name     string  `json:"name"`
	Quantity float32 `json:"count"`
	Price    float32 `json:"price"`
	Amount   float32 `json:"summa"`
}

type InvoiceCard struct {
	Name string `json:"name"`
	Rfid []uint `json:"card"`
}

var (
	InvoiceInState = transition.New(&InvoiceIn{})
	// InvoiceInItemState = transition.New(&InvoiceInItem{})
)

func init() {
	InvoiceInState.Initial("draft")
}
