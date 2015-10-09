package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/grengojbo/qor-example/config"
	"github.com/grengojbo/qor-example/config/admin"
	_ "github.com/grengojbo/qor-example/db/migrations"
)

var (
	Version   = "0.1.0"
	BuildTime = "2015-09-20 UTC"
	GitHash   = "c00"
)

func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

func main() {
	fmt.Printf("App Version: %s\n", Version)
	fmt.Printf("Build Time: %s\n", BuildTime)
	fmt.Printf("Git Commit Hash: %s\n", GitHash)
	fmt.Printf("Listening on: %v\n", config.Config.Port)
	ConfigRuntime()
	mux := http.NewServeMux()
	admin.Admin.MountTo("/admin", mux)

	r := gin.Default()
	r.LoadHTMLGlob("app/views/*.tmpl")
	for _, path := range []string{"system", "javascripts", "stylesheets", "images"} {
		// mux.Handle(fmt.Sprintf("/%s/", path), http.FileServer(http.Dir("public")))
		r.Static(fmt.Sprintf("/%s", path), fmt.Sprintf("public/%s", path))
	}

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.tmpl", gin.H{
			// "roomid":    roomid,
			// "nick":      nick,
			"timestamp": time.Now().Unix(),
		})
	})
	// r.POST("/login", func(c *gin.Context) {
	// 	var form models.User
	// 	// This will infer what binder to use depending on the content-type header.
	// 	if c.Bind(&form) == nil {
	// 		if form.User == "manu" && form.Password == "123" {
	// 			c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	// 		} else {
	// 			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
	// 		}
	// 	}
	// })
	r.GET("/logout", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/login")
	})

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/admin")
	})

	r.Any("/admin/*w", gin.WrapH(mux))
	r.Run(fmt.Sprintf(":%d", config.Config.Port))
}
