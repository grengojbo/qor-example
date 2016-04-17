package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/media_library"
)

type ExchangeSkeleton struct {
	gorm.Model
	Name     string
	Format   string
	Skeleton string
	Example  media_library.FileSystem
}
