package models

import (
	"github.com/jinzhu/gorm"
	// "github.com/qor/qor/media_library"
)

type User struct {
	gorm.Model
	Name      string
	Password  string
	FirstName string
	LastNname string
	Gender    string
	Role      string
	Email     []Email
	Phone     []Phone
	Social    []Social
	Languages []Language `gorm:"many2many:user_languages;"`
	Addresses []Address
	Comment   string
	// Location  string
	// Avatar    media_library.FileSystem
}

type Language struct {
	gorm.Model
	Name string
}

func (user User) DisplayName() string {
	return user.Name
}

func (user User) AvailableLocales() []string {
	return []string{"en-US", "zh-CN", "ru-RU"}
}
