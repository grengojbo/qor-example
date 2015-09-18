package models

import (
	"github.com/jinzhu/gorm"
	// "github.com/qor/qor/media_library"
)

type User struct {
	gorm.Model
	Email     string
	Password  string
	Name      string
	FirstName string
	LastNname string
	Gender    string
	Role      string
	Languages []Language `gorm:"many2many:user_languages;"`
	// Avatar    media_library.FileSystem
	Addresses []Address
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
