package models

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

type BannerShow struct {
	gorm.Model
	// BannerID         sql.NullInt64
	StoreID          uint
	SesUuid          string         `json:"-"`
	IsBot            bool           `json:"-"`
	UserMac          sql.NullString `sql:"size:100"`
	UserIp           string         `sql:"size:15"`
	IPv4             bool           `gorm:"column:ipv4" json:"ip_v4"`
	UserAgent        string         `sql:"type:char(2)"`
	AcceptLanguage   string         `sql:"type:char(2)"`
	Referrer         sql.NullString
	UaBrowserFamily  sql.NullString `sql:"size:20"`
	UaBrowserVersion sql.NullInt64  `json:"browser_version"`
	UaOsFamily       sql.NullString `sql:"size:20"`
	UaOsVersion      sql.NullString `sql:"size:10"`
	UaDeviceFamily   sql.NullString `sql:"size:20"`
	ShowYear         int16          `json:"-"`
	ShowMonth        int8           `json:"-"`
	ShowDay          int8           `json:"-"`
	ShowHour         int8           `json:"-"`
	ShowMinute       int8           `json:"-"`
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
