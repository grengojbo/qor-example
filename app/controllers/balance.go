package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/app/utils"
	"github.com/qor/qor-example/db"
	// "github.com/gin-gonic/contrib/sessions"
)

// GET http://localhost:7000/api/v1/balances
func BalanceApiIndex(ctx *gin.Context) {
	var balances []models.Balance
	var api []models.BalanceApi
	var count int
	var currentBtn int
	var next string
	var priv string
	offset := 0
	status := "error"
	message := "Not Found"

	// col := ctx.DefaultQuery("col", "none")
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "1000"))
	s := ctx.Query("offset")
	if len(s) > 0 {
		offset, _ = strconv.Atoi(s)
	}
	currentBtn = offset
	offset = offset * limit

	db.DB.Find(&balances).Count(&count)
	db.DB.Find(&balances).Limit(limit).Offset(offset).Count(&count)
	if len(balances) > 0 {
		for _, item := range balances {
			status = "success"
			message = ""
			line := models.BalanceApi{
				ID:        item.ID,
				Code:      item.Code,
				ProductID: item.ProductID,
				Count:     item.Count,
				Price:     item.Price,
				UserID:    item.UserID,
				StoreID:   item.StoreID,
				State:     item.State,
				Comment:   item.Comment,
			}
			api = append(api, line)
		}
	}
	total := utils.GetNumberOfButtonsForPagination(count, limit)
	if currentBtn > 0 {
		priv = fmt.Sprintf("?limit=%d&offset=%d", limit, currentBtn-1)
	}
	if currentBtn < (total - 1) {
		next = fmt.Sprintf("?limit=%d&offset=%d", limit, currentBtn+1)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  status,
		"message": message,
		"data":    &api,
		"count":   count,
		"total":   total,
		"current": currentBtn + 1,
		"next":    next,
		"priv":    priv,
	})
}

// GET http://localhost:7000/api/v1/balances/0
func BalanceApiShow(ctx *gin.Context) {
	var (
		balance models.Balance
		api     models.BalanceApi
	)
	id, err := strconv.Atoi(ctx.Param("id"))

	if err == nil {
		if !db.DB.Where("id = ?", id).First(&balance).RecordNotFound() {
			api.ID = balance.ID
			api.Code = balance.Code
			api.ProductID = balance.ProductID
			api.Count = balance.Count
			api.Price = balance.Price
			api.UserID = balance.UserID
			api.StoreID = balance.StoreID
			api.State = balance.State
			api.Comment = balance.Comment
		}
	}
	ctx.JSON(http.StatusOK, &api)
}

// POST curl -i -X POST -H "Content-Type: application/json"
func BalanceNewApi(ctx *gin.Context) {
	var balance models.BalanceApi

	if err := ctx.BindJSON(&balance); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": fmt.Sprintf("Bad request : %v", err)})
	} else {
		bal := models.Balance{
			Code:         balance.Code,
			ProductID:    balance.ProductID,
			Count:        balance.Count,
			Price:        balance.Price,
			UserID:       balance.UserID,
			StoreID:      balance.StoreID,
			SubscribedAt: time.Now(),
			State:        "draft",
			Comment:      balance.Comment,
		}
		if err := db.DB.Create(&bal).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": fmt.Sprintf("Is not save : %v", err)})
		} else {
			balance.ID = bal.ID
			ctx.JSON(http.StatusOK, &balance)
		}
	}
}
