package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/qor/qor-example/db"
	"github.com/qor/qor-example/app/models"
	// "github.com/qor/qor-example/app/utils"
	// "github.com/jinzhu/gorm"
	// "github.com/qor/transition"
	// "github.com/qor/validations"
)

func OrganizationIndex(ctx *gin.Context) {
	// var self models.OrganizationApi

	status := "error"
	message := "Not Found"
	statusCode := http.StatusNotFound
	res := models.ListOrganization()
	if len(res) > 0 {
		statusCode = http.StatusOK
		status = "success"
		ctx.JSON(statusCode, gin.H{
			"status": status,
			"data":   &res,
		})

	} else {
		ctx.JSON(statusCode, gin.H{
			"status":  status,
			"message": message,
		})
	}
}

func OrganizationAddApi(ctx *gin.Context) {
	var self models.Organization

	status := "error"
	message := "Not Found"
	statusCode := http.StatusNotFound

	if err := ctx.BindJSON(&self); err != nil {
		message = fmt.Sprintf("Bad request : %v", err)
		statusCode = http.StatusBadRequest
		ctx.JSON(statusCode, gin.H{"status": status, "message": message})
	} else {
		ctx.JSON(statusCode, gin.H{
			"status":  status,
			"message": message,
		})
	}
}
