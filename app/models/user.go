package models

import (
	"github.com/jinzhu/gorm"
	// "github.com/qor/qor/media_library"
)

type User struct {
	gorm.Model
	// Email     string `sql:"type:varchar(75)" json:"email"`
	Name      string `gorm:"column:username" sql:"type:varchar(30);unique_index" json:"username"`
	Password  string `sql:"type:varchar(128)" json:"-"`
	IsActive  bool   `gorm:"column:is_active"json:"active"`
	FirstName string `sql:"type:varchar(30)" json:"first_name"`
	LastName  string `sql:"type:varchar(30)" json:"last_name"`
	Gender    string
	Role      string
	// Email     []Email
	// Phone     []Phone
	// Social    []Social
	Languages []Language `gorm:"many2many:user_languages;"`
	Addresses []Address
	Comment   string
	// Location  string
	// Avatar    media_library.FileSystem
}

// func (user User) TableName() string {
// 	return "auth_user"
// }

type Language struct {
	gorm.Model
	Name string
}

func (user User) DisplayName() string {
	return user.Name
}

func (user User) AvailableLocales() []string {
	return []string{"en-US", "uk-UA", "ru-RU"}
}
