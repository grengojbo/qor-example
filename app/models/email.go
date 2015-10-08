package models

import "github.com/jinzhu/gorm"

type Email struct {
	gorm.Model
	UserID uint
	Email  string
}

// func (this Phone) Stringify() string {
// 	return fmt.Sprintf("%v, %v, %v", address.Address2, address.Address1, address.City)
// }
