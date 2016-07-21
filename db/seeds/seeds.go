package seeds

import (
	"log"
	"math/rand"
	"path/filepath"
	"time"

	"github.com/azumads/faker"
	"github.com/jinzhu/configor"
	"github.com/qor/publish"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/db"
)

var Fake *faker.Faker

var Seeds = struct {
	Units []struct {
		Code     string
		Name     string
		FullName string
	}
	Roles []struct {
		Name string
	}
	Languages []struct {
		Name string
		Code string
	}
	Categories []struct {
		Name string
	}
	Colors []struct {
		Name string
		Code string
	}
	Sizes []struct {
		Name string
		Code string
	}
	Products []struct {
		CategoryName string
		Collections  []struct {
			Name string
		}
		Name          string
		ZhName        string
		NameWithSlug  string
		Code          string
		Price         float32
		Description   string
		ZhDescription string
		MadeCountry   string
		ZhMadeCountry string

		ColorVariations []struct {
			ColorName string
			ColorCode string
			Images    []struct {
				URL string
			}
		}
		SizeVariations []struct {
			SizeName string
		}
	}
	Stores []struct {
		Name      string
		Phone     string
		Email     string
		Country   string
		Zip       string
		City      string
		Region    string
		Address   string
		Latitude  float64
		Longitude float64
	}
	Setting struct {
		ShippingFee     uint
		GiftWrappingFee uint
		CODFee          uint `gorm:"column:cod_fee"`
		TaxRate         int
		Address         string
		City            string
		Region          string
		Country         string
		Zip             string
		Latitude        float64
		Longitude       float64
	}
	Seo struct {
		SiteName    string
		DefaultPage struct {
			Title       string
			Description string
			Keywords    string
		}
		HomePage struct {
			Title       string
			Description string
			Keywords    string
		}
		ProductPage struct {
			Title       string
			Description string
			Keywords    string
		}
	}
	Enterprises []struct {
		Name           string
		Begins         string
		Expires        string
		RequiresCoupon bool
		Unique         bool

		Coupons []struct {
			Code string
		}
		Rules []struct {
			Kind  string
			Value string
		}
		Benefits []struct {
			Kind  string
			Value string
		}
	}
}{}

func init() {
	Fake, _ = faker.New("en")
	Fake.Rand = rand.New(rand.NewSource(42))
	rand.Seed(time.Now().UnixNano())

	filepaths, _ := filepath.Glob("db/seeds/data/*.yml")
	if err := configor.Load(&Seeds, filepaths...); err != nil {
		panic(err)
	}
}

func TruncateTables(tables ...interface{}) {
	for _, table := range tables {
		if err := db.DB.DropTableIfExists(table).Error; err != nil {
			panic(err)
		}
		if err := db.Publish.DraftDB().DropTableIfExists(table).Error; err != nil {
			panic(err)
		}
		db.DB.AutoMigrate(table)
		if publish.IsPublishableModel(table) {
			db.Publish.AutoMigrate(table)
		}
	}
}

func CreateUnits() {
	for _, c := range Seeds.Units {
		// fmt.Println(c)
		u := models.Unit{}
		u.Name = c.Name
		u.Code = c.Code
		u.FullName = c.FullName
		if err := db.DB.Where(models.Unit{Name: c.Name}).Assign(u).FirstOrCreate(&u).Error; err != nil {
			log.Fatalf("create unit (%v) failure, got err %v", u.Name, err)
		}
	}
}

func CreateRoles() {
	for _, c := range Seeds.Roles {
		role := models.Role{}
		role.Name = c.Name
		if err := db.DB.Where(models.Role{Name: c.Name}).FirstOrCreate(&role).Error; err != nil {
			log.Fatalf("create role (%v) failure, got err %v", role, err)
		}
	}
}

func CreateLanguages() {
	for _, c := range Seeds.Languages {
		language := models.Language{}
		language.Name = c.Name
		language.Code = c.Code
		if err := db.DB.Where(models.Language{Name: c.Name}).Assign(language).FirstOrCreate(&language).Error; err != nil {
			log.Fatalf("create language (%v) failure, got err %v", language, err)
		}
	}
}

func CreateCategories() {
	for _, c := range Seeds.Categories {
		category := models.Category{}
		category.Name = c.Name
		if err := db.DB.Create(&category).Error; err != nil {
			log.Fatalf("create category (%v) failure, got err %v", category, err)
		}
	}
}

func CreateSetting() {
	setting := models.Setting{}

	setting.ShippingFee = Seeds.Setting.ShippingFee
	setting.GiftWrappingFee = Seeds.Setting.GiftWrappingFee
	setting.CODFee = Seeds.Setting.CODFee
	setting.TaxRate = Seeds.Setting.TaxRate
	setting.Address = Seeds.Setting.Address
	setting.Region = Seeds.Setting.Region
	setting.City = Seeds.Setting.City
	setting.Country = Seeds.Setting.Country
	setting.Zip = Seeds.Setting.Zip
	setting.Latitude = Seeds.Setting.Latitude
	setting.Longitude = Seeds.Setting.Longitude

	if err := db.DB.Create(&setting).Error; err != nil {
		log.Fatalf("create setting (%v) failure, got err %v", setting, err)
	}
}
