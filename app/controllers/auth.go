package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/grengojbo/gotools"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/admin"
	"github.com/qor/qor-example/db"
)

//GET Login
func LoginForm(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Set("lastLogin", time.Now().Unix())
	session.Delete("_auth_user_id")
	session.Set("ip", ctx.ClientIP())
	session.Save()
	// ctx.HTML(200, "login.tmpl", gin.H{
	config.View.Layout("main").Funcs(I18nFuncMap(ctx)).Execute(
		"login_jbo",
		gin.H{
			"title":     config.Config.SiteName,
			"timestamp": time.Now().Unix(),
		},
		ctx.Request,
		ctx.Writer,
	)
}

// POST Login
func Login(ctx *gin.Context) {
	var login admin.Auth
	session := sessions.Default(ctx)
	session.Delete("_auth_user_id")
	session.Save()
	if ctx.BindJSON(&login) == nil {
		if ok, user := login.GetUser(); ok != false {
			if err := gotools.VerifyPassword(user.Password, login.Password); err != nil {
				session.Set("lastLogin", time.Now().Unix())
				session.Save()
				ctx.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized", "message": "User unauthorized"})
			} else {
				session.Set("lastLogin", time.Now().Unix())
				session.Set("_auth_user_id", user.ID)
				session.Save()

				t := time.Now()
				login := models.LogLogin{
					ClietIp:   ctx.ClientIP(),
					UserID:    user.ID,
					InOut:     "in",
					LoginedAt: &t,
					Device:    "web",
				}
				if err := db.DB.Create(&login).Error; err != nil {
					fmt.Println(err)
				}

				ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Ok"})
			}
		} else {
			session.Set("lastLogin", time.Now().Unix())
			session.Save()
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized", "message": "User unauthorized"})
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Bad request"})
	}
}

// curl -i -X POST -H "Content-Type: application/json" -d "{ \"username\": \"pos0001\", \"password\": \"123456\", \"metod\": \"password\" }" http://localhost:7000/api/v1/auth
func LoginApi(ctx *gin.Context) {
	var auth models.Auth
	var currentUser models.User
	if ctx.BindJSON(&auth) == nil {
		if auth.Metod == "password" {
			if !db.DB.Where("password = ?", auth.Password).First(&currentUser).RecordNotFound() {
				t := time.Now()
				login := models.LogLogin{
					ClietIp:   ctx.ClientIP(),
					UserID:    currentUser.ID,
					User:      currentUser,
					InOut:     "in",
					LoginedAt: &t,
					Device:    "terminal",
				}
				if err := db.DB.Create(&login).Error; err != nil {
					fmt.Println(err)
				}
				user := models.UserApi{
					ID:     currentUser.ID,
					Name:   currentUser.Name,
					Email:  currentUser.Email,
					Gender: currentUser.Gender,
					Role:   currentUser.Role,
					Token:  "",
				}
				ctx.JSON(http.StatusOK, &user)
			} else {
				ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "User not found"})
			}
		} else if auth.Metod == "web" {
			if !db.DB.Where("name = ? OR email = ?", auth.User, auth.User).First(&currentUser).RecordNotFound() {
				if err := gotools.VerifyPassword(currentUser.Password, auth.Password); err != nil {
					// session.Set("lastLogin", time.Now().Unix())
					// session.Save()
					ctx.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized", "message": "User unauthorized"})
				} else {
					// session.Set("lastLogin", time.Now().Unix())
					// session.Set("_auth_user_id", user.ID)
					// session.Save()
					// ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Ok"})
					t := time.Now()
					login := models.LogLogin{
						ClietIp:   ctx.ClientIP(),
						UserID:    currentUser.ID,
						User:      currentUser,
						InOut:     "in",
						LoginedAt: &t,
						Device:    "terminal",
					}
					if err := db.DB.Create(&login).Error; err != nil {
						fmt.Println(err)
					}
					user := models.UserApi{
						ID:     currentUser.ID,
						Name:   currentUser.Name,
						Email:  currentUser.Email,
						Gender: currentUser.Gender,
						Role:   currentUser.Role,
						Token:  "",
					}
					ctx.JSON(http.StatusOK, &user)
				}
			}
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Bad request metod"})
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Bad request"})
	}
}

func LoginJWT(ctx *gin.Context) {
	// Create the token
	token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
	// Set some claims
	// token.Claims["ID"] = "Christopher"
	// token.Claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(config.Config.Token))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not generate token"})
	}
	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Delete("_auth_user_id")
	if err := session.Save(); err != nil {
		log.Println("[ERROR]", err)
	}
	log.Println("[LOGOUT]")
	// log.Println(session)
	ctx.Redirect(http.StatusMovedPermanently, "/login")
}

func LogoutApi(ctx *gin.Context) {
	var currentUser models.User
	if !db.DB.Where("id = ?", ctx.Param("id")).First(&currentUser).RecordNotFound() {
		t := time.Now()
		login := models.LogLogin{
			ClietIp:   ctx.ClientIP(),
			UserID:    currentUser.ID,
			User:      currentUser,
			InOut:     "out",
			LoginedAt: &t,
			Device:    "terminal",
		}
		if err := db.DB.Create(&login).Error; err != nil {
			fmt.Println(err)
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": fmt.Sprintf("User %s exit.", currentUser.Name)})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "User not found"})
	}
}
