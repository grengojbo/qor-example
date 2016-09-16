package main

import (
	"fmt"
	"net/http"

	//go:generate go-bindata -nomemcopy ../qor/admin/views/...
	// "github.com/gin-gonic/contrib/sessions"
	// "github.com/grengojbo/gotools"
	"github.com/gin-gonic/gin"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/admin"
	"github.com/qor/qor-example/config/api"
	_ "github.com/qor/qor-example/config/i18n"
	"github.com/qor/qor-example/config/routes"
	"github.com/qor/qor-example/db/migrations"
)

var (
	Version   = "0.1.0"
	BuildTime = "2015-09-20 UTC"
	GitHash   = "c00"
)

func main() {
	conf := config.Config
	fmt.Printf("App Version: %s\n", Version)
	fmt.Printf("Build Time: %s\n", BuildTime)
	fmt.Printf("Git Commit Hash: %s\n", GitHash)
	fmt.Printf("Listening on: %v\n", conf.Port)

	if conf.Migration {
		migrations.Run()
	}

	mux := http.NewServeMux()
	// mux.Handle("/", routes.Router())
	// mux.Handle("/auth/", auth.Auth.NewRouter())
	admin.Admin.MountTo("/admin", mux)
	api.API.MountTo("/api", mux)
	config.Filebox.MountTo("/downloads", mux)

	for _, path := range []string{"system", "javascripts", "stylesheets", "images"} {
		mux.Handle(fmt.Sprintf("/%s/", path), http.FileServer(http.Dir("public")))
	}

	fmt.Printf("Listening on: %v\n", config.Config.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), mux); err != nil {
		panic(err)
	}
}
