package controllers_test

import (
	"encoding/json"
	"testing"

	"net/http"
	"net/http/httptest"

	_ "github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/routes"
	. "github.com/smartystreets/goconvey/convey"
)

// var r *gin.Engine

// func init() {
// r = routes.Router()
// }

func TestCategory(t *testing.T) {
	// t.Parallel()
	Convey("Test Category controller", t, func() {
		r := routes.Router()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/category", nil)
		// req.Header.Set("Content-Type", "application/json")
		// req.Header.Set("Locale", "ru-RU")
		// req.Header.Set("Accept-Language", "ru-RU,ru;q=0.8,en-US;q=0.6,en;q=0.4,uk;q=0.2")
		r.ServeHTTP(w, req)

		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
			_, err := json.Unmarshal(w.Body)
			So(err, ShouldBeNil)
		})
	})
}
