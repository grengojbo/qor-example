package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qor/qor-example/config"
)

func IpMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// fmt.Println("Remote IP: ", ctx.ClientIP())
		if config.Config.Allow != "any" && config.Config.Allow != ctx.ClientIP() {
			ctx.JSON(http.StatusForbidden, gin.H{"error": fmt.Sprintf("Forbidden: %v", ctx.ClientIP())})
			ctx.Abort()
			return
		} else {
			ctx.Next()
		}
	}
}
