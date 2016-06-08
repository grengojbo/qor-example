package config

import (
	"log"
	"os"

	"github.com/jinzhu/configor"
	// "github.com/qor/i18n"
	"github.com/qor/render"
)

type SMTPConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Site     string
}

var Config = struct {
	SiteName  string `default:"Qor DEMO"`
	ApiUrl    string `default:"http://localhost:7000/api/"`
	Token     string `default:"mysupersecretpassword"`
	Migration bool   `default:"true"`
	Env       string `env:"ENV" default:"local"`
	Port      uint   `default:"7000" env:"PORT"`
	DB        struct {
		Name     string `default:"qor_example"`
		Adapter  string `default:"mysql"`
		User     string
		Password string
		Host     string `default:"localhost"`
		Port     uint   `default:"3306"`
		Debug    bool   `default:"false"`
	}
	SMTP SMTPConfig
	Log  struct {
		FileName string
		Maxdays  int `default:"30"`
	}
	Redis struct {
		Host     string `default:"localhost"`
		Port     uint   `default:"6379"`
		Protocol string `default:"tcp"`
		Password string
	}
	Session struct {
		Name    string `default:"sessionid"`
		Adapter string `default:"cookie"`
	}
	// I18n           *i18n.I18n
	Locale         string `default:"en-US"`
	Secret         string `default:"secret"`
	Limit          int    `default:"5"`
	PasswordLength int    `default:"6"`
}{}

var (
	Root       = os.Getenv("GOPATH") + "/src/github.com/qor/qor-example"
	FileConfig = os.Getenv("QORROOT") + "/config/database.yml"
	FileSMTP   = os.Getenv("QORROOT") + "/config/smtp.yml"
	View       *render.Render
)

// Set environment variable config path -> export QORCONFIG=/etc/qor/production.yml
func init() {
	if file := os.Getenv("QORCONFIG"); len(file) > 0 {
		FileConfig = file
		log.Printf("Set config file: %s\n", FileConfig)
	}
	if file := os.Getenv("QORSMTP"); len(file) > 0 {
		FileSMTP = file
		log.Printf("Set config file: %s\n", FileSMTP)
	}
	if rootPath := os.Getenv("QORROOT"); len(rootPath) > 0 {
		Root = rootPath
		log.Printf("Set ROOT path: %s\n", FileConfig)
	}
	if err := configor.Load(&Config, FileConfig, FileSMTP); err != nil {
		panic(err)
	}

	View = render.New()
}

func (s SMTPConfig) HostWithPort() string {
	return s.Host + ":" + s.Port
}
