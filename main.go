package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/grengojbo/qor-example/config"
	"github.com/grengojbo/qor-example/config/admin"
	_ "github.com/grengojbo/qor-example/db/migrations"
	// "github.com/grengojbo/qor-example/db"
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
	conf := config.Config
	fmt.Printf("App Version: %s\n", Version)
	fmt.Printf("Build Time: %s\n", BuildTime)
	fmt.Printf("Git Commit Hash: %s\n", GitHash)
	fmt.Printf("Listening on: %v\n", conf.Port)
	// fmt.Printf("Secret: %v\n", admin.PasswordBcrypt("admin"))
	ConfigRuntime()
	mux := http.NewServeMux()
	admin.Admin.MountTo("/admin", mux)

	r := gin.Default()
	if conf.Session.Adapter == "redis" {
		fmt.Println("Connect to redis")
		store, err := sessions.NewRedisStore(10, conf.Redis.Protocol, fmt.Sprintf("%v:%v", conf.Redis.Host, conf.Redis.Port), "", []byte(conf.Secret))
		if err != nil {
			panic(err)
		}
		fmt.Println("Start init session")
		r.Use(sessions.Sessions(conf.Session.Name, store))
		fmt.Println("Finish init session")
	}
	r.LoadHTMLGlob("app/views/*.tmpl")
	for _, path := range []string{"system", "javascripts", "stylesheets", "images"} {
		// mux.Handle(fmt.Sprintf("/%s/", path), http.FileServer(http.Dir("public")))
		r.Static(fmt.Sprintf("/%s", path), fmt.Sprintf("public/%s", path))
	}

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.GET("/login", func(c *gin.Context) {
		// var cnt int
		session := sessions.Default(c)
		// if s, ok := db.DB.Get();
		count := session.Get("count")
		if count == nil {
			count = 0
		} else {
			cnt := count.(int)
			count = cnt + 1
		}
		session.Set("count", count)
		session.Set("lastLogin", time.Now().Unix())
		session.Set("ip", c.ClientIP())
		session.Save()
		c.HTML(200, "login.tmpl", gin.H{
			"title": admin.Admin.SiteName,
			// "nick":      nick,
			"timestamp": time.Now().Unix(),
		})
	})
	r.POST("/login", func(c *gin.Context) {
		var login admin.Auth
		session := sessions.Default(c)
		count := session.Get("count")
		if count == nil {
			count = 0
		} else {
			cnt := count.(int)
			count = cnt + 1
		}
		session.Set("count", count)
		if count.(int) > conf.Limit {
			fmt.Printf("login limit count: %v\n", count)
			session.Set("lastLogin", time.Now().Unix())
			session.Save()
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Bad request"})
		} else {
			if c.BindJSON(&login) == nil {
				if ok, user := login.GetUser(); ok != false {
					if err := admin.VerifyPassword(user.Password, login.Password); err != nil {
						session.Set("lastLogin", time.Now().Unix())
						session.Save()
						c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized", "message": "User unauthorized"})
					} else {
						session.Set("count", 0)
						session.Set("_auth_user_id", user.ID)
						session.Set("user", user)
						// c.Request.Cookie("userid", user.ID, "", "",)
						session.Save()
						// u := session.Get("user")
						// fmt.Printf("%v\n", u)
						// fmt.Println("-------------------------")
						// fmt.Printf("%v\n", u.Email)
						// fmt.Println("-------------------------")
						c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Ok"})
					}
				} else {
					session.Set("lastLogin", time.Now().Unix())
					session.Save()
					c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized", "message": "User unauthorized"})
				}
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Bad request"})
			}
		}
	})
	r.GET("/logout", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Save()
		c.Redirect(http.StatusMovedPermanently, "/login")
	})

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/admin")
	})

	r.Any("/admin/*w", gin.WrapH(mux))
	r.Run(fmt.Sprintf(":%d", config.Config.Port))
}
