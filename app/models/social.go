package models

import "github.com/jinzhu/gorm"

type Social struct {
	gorm.Model
	UserID uint
	Name   string
}

// func (this Phone) Stringify() string {
// 	return fmt.Sprintf("%v, %v, %v", address.Address2, address.Address1, address.City)
// }
