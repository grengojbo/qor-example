package models_test

import (
	"testing"

	"github.com/qor/qor-example/app/models"
	_ "github.com/qor/qor-example/config"
	"github.com/qor/qor-example/db/migrations"
	. "github.com/smartystreets/goconvey/convey"
)

func TestModelsCategory(t *testing.T) {
	// t.Parallel()
	migrations.Run()
	Convey("Test Category model", t, func() {

		Convey("List All Category", func() {
			res := models.GetAllCategory()
			So(res, ShouldHaveLength, 3)
		})

	})
}
