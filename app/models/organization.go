package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/location"
	"github.com/qor/media_library"
	"github.com/qor/qor-example/db"
)

type OrganizationApi struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Organization struct {
	gorm.Model
	Name     string `gorm:"column:name" sql:"type:varchar(30);unique_index" json:"name"`
	Enabled  bool   `sql:"default:false" json:"-"`
	Director string `json:"director"`
	Email    string `sql:"type:varchar(75)" json:"email"`
	// Phone          []Phone
	Logo           media_library.FileSystem
	Edrpou         uint
	Ipn            uint
	SvidPlanNaloga uint
	Score          uint64
	Bank           string
	Mfo            uint
	// Addresses      []Address
	Comment string
	location.Location
}

func (organization Organization) DisplayName() string {
	return organization.Name
}

func ListOrganization() (res []OrganizationApi) {
	self := []Organization{}
	api := []OrganizationApi{}
	if err := db.DB.Where(&Organization{Enabled: true}).Find(&self).Error; err != nil {
		return api
	}
	for _, row := range self {
		item := OrganizationApi{
			ID:   row.ID,
			Name: row.Name,
		}
		api = append(api, item)
	}
	return api

}
