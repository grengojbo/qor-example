package controllers_test

import (
	"testing"

	"net/http"
	"net/http/httptest"

	r_t "github.com/qor/qor-example/config/routes"
	. "github.com/smartystreets/goconvey/convey"
)

// var r *gin.Engine

// func init() {
// r = routes.Router()
// }

func TestCategory(t *testing.T) {
	t.Parallel()
	Convey("Test Category controller", t, func() {
		r := r_t.Router()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/category", nil)
		r.ServeHTTP(w, req)

		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})
}
