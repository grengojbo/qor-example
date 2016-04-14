package models

import (
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
	Invoices       []InvoiceOutItem `json:"invoices"`
	transition.Transition
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
