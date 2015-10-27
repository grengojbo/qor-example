package models

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
)

type BannerShow struct {
	gorm.Model
	ShowDate         time.Time      `sql:"type:date not null;index" json:"show_date"`
	ShowTime         time.Time      `sql:"type:time with time zone not null" json:"show_time"`
	BannerID         int            `json:"banner_id"`
	CampaignID       int            `json:"campaign_id"`
	StoreID          int            `json:"region_id"`
	IsBot            bool           `json:"-"`
	SesUuid          string         `json:"-"`
	UserMac          sql.NullString `sql:"type:char(23)"`
	UserIp           string         `sql:"size:15"`
	IPv4             bool           `gorm:"column:ipv4" json:"ip_v4"`
	AcceptLanguage   string         `sql:"type:char(2)"`
	UaBrowserFamily  sql.NullString `sql:"size:20"`
	UaBrowserVersion int16          `json:"browser_version"`
	UaOsFamily       sql.NullString `sql:"size:100"`
	UaOsVersion      sql.NullString `sql:"size:10"`
	UaDeviceFamily   sql.NullString `sql:"size:20"`
	IsMobile         bool           `gorm:"column:is_mobile" json:"is_mobile"`
	IsTable          bool           `gorm:"column:is_table" json:"is_table"`
	IsDesktop        bool           `gorm:"column:is_desktop" json:"is_desktop"`
	// ShowYear         int16          `json:"-"`
	// ShowMonth        int8           `json:"-"`
	// ShowDay          int8           `json:"-"`
	// ShowHour         int8           `json:"-"`
	// ShowMinute       int8           `json:"-"`
	Referrer  sql.NullString
	UserAgent string
	// sorting.SortingDESC
}

// type BannerShow struct {
// 	gorm.Model
// 	BannerID         sql.NullInt64
// 	UserID           sql.NullInt64
// 	ZoneID           sql.NullInt64
// 	ses              string
// 	UserMac          sql.NullString `sql:"size:100"`
// 	UserIp           string         `sql:"size:15"`
// 	UserAgent        string         `sql:"size:1000"`
// 	Referrer         string
// 	UaBrowserFamily  sql.NullString `sql:"size:20"`
// 	UaBrowserVersion sql.NullString `sql:"size:10"`
// 	UaOsFamily       sql.NullString `sql:"size:20"`
// 	UaOsVersion      sql.NullString `sql:"size:10"`
// 	UaDeviceFamily   sql.NullString `sql:"size:20"`
// }
