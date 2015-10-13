package migrations

import (
	"log"

	"github.com/grengojbo/qor-example/app/models"
	"github.com/grengojbo/qor-example/db"
	"github.com/qor/qor/admin"
	"github.com/qor/qor/publish"
)

var Admin *admin.Admin

func init() {
	log.Println("Start migration ...")
	AutoMigrate(&admin.AssetManager{})

	AutoMigrate(&models.Product{}, &models.ColorVariation{}, &models.ColorVariationImage{}, &models.SizeVariation{})
	AutoMigrate(&models.Color{}, &models.Size{}, &models.Category{}, &models.Collection{})

	AutoMigrate(&models.Address{})

	AutoMigrate(&models.Order{}, &models.OrderItem{})

	AutoMigrate(&models.Store{})

	AutoMigrate(&models.Newsletter{})

	AutoMigrate(&models.Setting{})

}

func AutoMigrate(values ...interface{}) {
	for _, value := range values {
		db.DB.AutoMigrate(value)

		if publish.IsPublishableModel(value) {
			db.Publish.AutoMigrate(value)
		}
	}
}
