package models

import "github.com/jinzhu/gorm"

type Phone struct {
	gorm.Model
	UserID uint
	Phone  string
}

// func (this Phone) Stringify() string {
// 	return fmt.Sprintf("%v, %v, %v", address.Address2, address.Address1, address.City)
// }
