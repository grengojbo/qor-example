package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/apertoire/mlog"
	"github.com/gin-gonic/gin"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/admin"
	"github.com/qor/qor-example/db"
	"github.com/qor/seo"
)

// GET: http://localhost:7000/api/v1/products
func ProductApiIndex(ctx *gin.Context) {
	mlog.Start(mlog.LevelTrace, "")
	var products []models.Product
	var count int
	var currentBtn int
	var next string
	var priv string
	offset := 0
	acceptLanguage := ctx.Request.Header.Get("Accept-Language")[0:2]
	locale := ctx.Request.Header.Get("Locale")
	if len(locale) == 0 {
		locale = config.Config.Locale
	}
	if err := db.DB.Set("l10n:locale", locale).Find(&products).Count(&count).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Bad request"})
	}
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "1000"))
	s := ctx.Query("offset")
	if len(s) > 0 {
		offset, _ = strconv.Atoi(s)
	}
	currentBtn = offset
	offset = offset * limit
	mlog.Trace("limit: %v offset: %v currentBtn: %v", limit, offset, currentBtn)
	mlog.Trace("acceptLanguage: %v, locale: %v", acceptLanguage, locale)
	if err := db.DB.Set("l10n:locale", locale).Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Bad request"})
	}
	total := getNumberOfButtonsForPagination(count, limit)
	if currentBtn > 0 {
		priv = fmt.Sprintf("?limit=%d&offset=%d", limit, currentBtn-1)
	}
	if currentBtn < (total - 1) {
		next = fmt.Sprintf("?limit=%d&offset=%d", limit, currentBtn+1)
	}
	// session := sessions.Default(ctx)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"data":    &products,
		"count":   count,
		"total":   total,
		"current": currentBtn + 1,
		"next":    next,
		"priv":    priv,
	})
}

func ProductIndex(ctx *gin.Context) {
	var (
		products   []models.Product
		seoSetting models.SEOSetting
	)

	db.DB.Limit(10).Find(&products)
	db.DB.First(&seoSetting)

	ctx.HTML(
		http.StatusOK,
		"product_index.tmpl",
		gin.H{
			"Products": products,
			"SeoTag":   seoSetting.DefaultPage.Render(seoSetting),
			"MicroSearch": seo.MicroSearch{
				URL:    "http://demo.getqor.com",
				Target: "http://demo.getqor.com/search?q=",
			}.Render(),
			"MicroContact": seo.MicroContact{
				URL:         "http://demo.getqor.com",
				Telephone:   "080-0012-3232",
				ContactType: "Customer Service",
			}.Render(),
		},
	)
}

func ProductShow(ctx *gin.Context) {
	var (
		product        models.Product
		colorVariation models.ColorVariation
		seoSetting     models.SEOSetting
		codes          = strings.Split(ctx.Param("code"), "_")
		productCode    = codes[0]
		colorCode      string
	)

	if len(codes) > 1 {
		colorCode = codes[1]
	}

	DB(ctx).Where(&models.Product{Code: productCode}).First(&product)
	DB(ctx).Preload("Images").Preload("Product").Preload("Color").Preload("SizeVariations.Size").Where(&models.ColorVariation{ProductID: product.ID, ColorCode: colorCode}).First(&colorVariation)
	DB(ctx).First(&seoSetting)

	config.View.Funcs(funcsMap(ctx)).Execute(
		"product_show",
		gin.H{
			"ActionBarTag":   admin.ActionBar.Render(ctx.Writer, ctx.Request),
			"Product":        product,
			"ColorVariation": colorVariation,
			"SeoTag":         seoSetting.ProductPage.Render(seoSetting, product),
			"MicroProduct": seo.MicroProduct{
				Name:        product.Name,
				Description: product.Description,
				BrandName:   product.Category.Name,
				SKU:         product.Code,
				Price:       float64(product.Price),
				Image:       colorVariation.MainImageUrl(),
			}.Render(),
			"CurrentUser":   CurrentUser(ctx),
			"CurrentLocale": CurrentLocale(ctx),
		},
		ctx.Request,
		ctx.Writer,
	)
}

func funcsMap(ctx *gin.Context) template.FuncMap {
	funcMaps := map[string]interface{}{
		"related_products": func(cv models.ColorVariation) []models.Product {
			var products []models.Product
			DB(ctx).Preload("ColorVariations").Preload("ColorVariations.Images").Limit(4).Find(&products, "id <> ?", cv.ProductID)
			return products
		},
		"other_also_bought": func(cv models.ColorVariation) []models.Product {
			var products []models.Product
			DB(ctx).Preload("ColorVariations").Preload("ColorVariations.Images").Order("id ASC").Limit(8).Find(&products, "id <> ?", cv.ProductID)
			return products
		},
	}
	for key, value := range I18nFuncMap(ctx) {
		funcMaps[key] = value
	}
	return funcMaps
}

func getNumberOfButtonsForPagination(TotalCount int, limit int) int {
	num := (int)(TotalCount / limit)
	if TotalCount%limit > 0 {
		num++
	}
	return num
}
