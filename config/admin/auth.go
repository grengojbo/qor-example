package admin

import (
	"github.com/grengojbo/qor-example/app/models"
	// "github.com/grengojbo/qor-example/db"
	// "github.com/jinzhu/gorm"
	"github.com/qor/qor"
	"github.com/qor/qor/admin"
)

// var DB *gorm.DB

type Auth struct{}

func (Auth) LoginURL(c *admin.Context) string {
	return "/login"
}

func (Auth) LogoutURL(c *admin.Context) string {
	return "/logout"
}

func (Auth) GetCurrentUser(c *admin.Context) qor.CurrentUser {
	var currentUser models.User
	if !c.GetDB().Where("name = ?", "admin").First(&currentUser).RecordNotFound() {
		return &currentUser
	}

	return nil
}
