package main

import (
	"fmt"
	"net"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/grengojbo/qor-example/config"
	"github.com/grengojbo/qor-example/config/admin"
	_ "github.com/grengojbo/qor-example/db/migrations"
	"github.com/mssola/user_agent"
	"github.com/nu7hatch/gouuid"
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
		store, err := sessions.NewRedisStore(10, conf.Redis.Protocol, fmt.Sprintf("%v:%v", conf.Redis.Host, conf.Redis.Port), "", []byte(conf.Secret))
		if err != nil {
			panic(err)
		}
		r.Use(sessions.Sessions(conf.Session.Name, store))
	} else if conf.Session.Adapter == "cookie" {
		store := sessions.NewCookieStore([]byte(conf.Secret))
		r.Use(sessions.Sessions(conf.Session.Name, store))
	}
	r.LoadHTMLGlob("app/views/*.tmpl")
	for _, path := range []string{"system", "javascripts", "stylesheets", "images"} {
		// mux.Handle(fmt.Sprintf("/%s/", path), http.FileServer(http.Dir("public")))
		r.Static(fmt.Sprintf("/%s", path), fmt.Sprintf("public/%s", path))
	}

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/show/ping", func(c *gin.Context) {
		// fmt.Printf("%v\n", c.Request)
		fmt.Printf("User-Agent: %s\n", c.Request.UserAgent())
		ua := user_agent.New(c.Request.UserAgent())
		fmt.Printf("%v\n", ua.Mobile())  // => false
		fmt.Printf("%v\n", ua.Bot())     // => false
		fmt.Printf("%v\n", ua.Mozilla()) // => "5.0"

		fmt.Printf("%v\n", ua.Platform()) // => "X11"
		fmt.Printf("%v\n", ua.OS())       // => "Linux x86_64"

		name, version := ua.Engine()
		fmt.Printf("%v\n", name)    // => "AppleWebKit"
		fmt.Printf("%v\n", version) // => "537.11"

		name, version = ua.Browser()
		fmt.Printf("Browser name:%v version: %v\n", name, version)
		if len(c.Request.Referer()) > 1 {
			fmt.Printf("Referrer: %s\n", c.Request.Referer())
		}
		// var ip string
		ipv4 := false
		ip, _, err := net.SplitHostPort(c.ClientIP())
		if err != nil {
			ip = "127.0.0.1"
			// } else {
			// 	ip = ipSrc.String()
		}
		// v := net.ParseIP(ip)
		if net.ParseIP(ip).To4() != nil {
			ipv4 = true
		}
		fmt.Printf("Remote IP: %s IPv4=%v\n", ip, ipv4)
		fmt.Printf("Accept-Language: %s\n", c.Request.Header.Get("Accept-Language")[0:2])
		u4, err := uuid.NewV4()
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		fmt.Println("uuid:", u4)
		// w.Header().Set("cache-control", "priviate, max-age=0, no-cache")
		c.Header("cache-control", "priviate, max-age=0, no-cache")
		c.Header("pragma", "no-cache")
		c.Header("expires", "-1")
		c.Header("Last-Modified", fmt.Sprintf("%v", time.Now().UTC()))
		// c.Header("Date", time.Now().UTC())
		// INSERT INTO "banner_shows" ( "ses_uuid", "store_id", "user_ip", "user_mac", "created_at", "updated_at", "show_year", "show_month", "show_day", "show_hour", "show_minute", "user_agent", "accept_language" )
		// VALUES ('df9e9ecf-b17c-4fe6-502e-aaddb55b961c', 8, '127.0.0.1:61526', '01:23:32:bb:63:12', NOW(), NOW(), Extract(YEAR from Now()), Extract(month from Now()), Extract(DAY from Now()), Extract(hour from Now()), Extract(minute from Now()), 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.101 Safari/537.36', 'ru');
		c.String(200, "pong")
	})

	r.GET("/login", func(c *gin.Context) {
		// var cnt int
		session := sessions.Default(c)
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
			"title":     admin.Admin.SiteName,
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
						// session.Set("user", user)
						session.Save()
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
