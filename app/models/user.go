package models

import (
	"database/sql"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/qor/media_library"
	"github.com/qor/qor-example/db"
)

type UserApi struct {
	ID      uint   `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Gender  string `json:"gender"`
	Role    string `json:"role"`
	Token   string `json:"token"`
	Enabled bool   `json:"enabled"`
}

type User struct {
	gorm.Model
	Email           string         `sql:"type:varchar(75)" json:"email"`
	Name            sql.NullString `gorm:"column:name" sql:"type:varchar(30)" json:"username"`
	Password        string         `sql:"type:varchar(128)" json:"-"`
	PasswordConfirm string         `gorm:"-" json:"-"`
	FirstName       string         `sql:"type:varchar(30)" json:"first_name"`
	LastName        string         `sql:"type:varchar(30)" json:"last_name"`
	OrganizationID  uint
	Organization    Organization
	Gender          string
	Role            string
	Languages       []Language `gorm:"many2many:user_languages;"`
	Addresses       []Address
	Comment         string
	Enabled         bool `sql:"default:true" json:"-"`
	Avatar          media_library.FileSystem

	// Confirm
	ConfirmToken string
	Confirmed    bool

	// Recover
	RecoverToken       string
	RecoverTokenExpiry *time.Time
}

// func (user User) TableName() string {
//  return "auth_user"
// }

func (user User) DisplayName() string {
	return user.Name.String
}

func (user User) AvailableLocales() []string {
	return []string{"en-US", "uk-UA", "ru-RU"}
}

// func (User) ViewableLocales() []string {
//   return []string{l10n.Global, "zh-CN", "JP", "EN", "DE"}
// }

// func (user User) EditableLocales() []string {
//   if user.role == "global_admin" {
//     return []string{l10n.Global, "zh-CN", "EN"}
//   } else {
//     return []string{"zh-CN", "EN"}
//   }
// }

// func (user User) Validate(db *gorm.DB) {
// 	if strings.TrimSpace(user.Password) == "" {
// 		db.AddError(validations.NewError(user, "Password", "Name can not be empty"))
// 	}

// 	if strings.TrimSpace(user.Name) == "" {
// 		db.AddError(validations.NewError(user, "Name", "Name can not be empty"))
// 	}
// }

type Language struct {
	gorm.Model
	Name string
	Code string
}

// User Role
type Role struct {
	gorm.Model
	Name string
}

func Roles() (results []string) {
	roleVariations := []Role{}
	if err := db.DB.Debug().Find(&roleVariations).Error; err != nil {
		log.Fatalf("query Role (%v) failure, got err %v", roleVariations, err)
		return results
	}
	for _, role := range roleVariations {
		results = append(results, role.Name)
	}
	return results
}
