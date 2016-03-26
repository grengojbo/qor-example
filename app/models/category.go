package models

import (
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/qor/l10n"
	"github.com/qor/publish"
	"github.com/qor/qor-example/db"
	"github.com/qor/sorting"
	"github.com/qor/validations"
)

type Category struct {
	gorm.Model
	l10n.Locale     `json:"-"`
	publish.Status  `json:"-"`
	sorting.Sorting `json:"-"`
	Name            string `json:"name"`
}

func (category Category) Validate(db *gorm.DB) {
	if strings.TrimSpace(category.Name) == "" {
		db.AddError(validations.NewError(category, "Name", "Name can not be empty"))
	}
}

func GetAllCategory() (res []Category) {
	c := []Category{}
	if err := db.DB.Find(&c).Error; err != nil {
		return c
	}
	return c
}
