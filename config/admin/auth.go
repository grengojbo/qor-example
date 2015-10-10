package admin

import (
	"fmt"

	"code.google.com/p/go.crypto/bcrypt"
	"github.com/grengojbo/qor-example/app/models"
	// "github.com/jinzhu/gorm"
	"github.com/grengojbo/qor-example/db"
	"github.com/qor/qor"
	"github.com/qor/qor/admin"
)

// var DB *gorm.DB

type Auth struct {
	User     string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

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

// Return bcrypt password
func PasswordBcrypt(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic("Permissions: bcrypt password hashing unsuccessful")
	}
	return string(hash)
}

func VerifyPassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (this *Auth) GetUser() (bool, *models.User) {
	var currentUser models.User
	if !db.DB.Where("name = ?", this.User).First(&currentUser).RecordNotFound() {
		fmt.Println("User:", this.Password)
		return true, &currentUser
	}
	return false, &currentUser
}
