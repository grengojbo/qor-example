package client

import (
	// "crypto/tls"
	// "fmt"
	// "io/ioutil"
	// "log"
	// "os"
	// "strconv"
	// "time"

	// _ "github.com/qor/qor-example/config"
	"fmt"

	"github.com/go-resty/resty"
	"github.com/jinzhu/gorm"
)

type Client struct {
	Url    string
	Debug  bool
	DB     *gorm.DB
	Config *Config
	Auth   AuthSuccess
}

type Config struct {
	Host     string
	Protocol string
	Token    string
}

type AuthSuccess struct {
	/* variables */
}
type AuthError struct {
	/* variables */
}

func (self *Client) Init() {
	self.Url = fmt.Sprintf("%s://%s/api/", self.Config.Protocol, self.Config.Host)
}

func (self *Client) LoadCategory() {
	resp, err := resty.R().Get("http://httpbin.org/get")

	fmt.Printf("\nError: %v", err)
	fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
	fmt.Printf("\nResponse Status: %v", resp.Status())
	fmt.Printf("\nResponse Body: %v", resp)
	fmt.Printf("\nResponse Time: %v", resp.Time())
	fmt.Printf("\nResponse Recevied At: %v", resp.ReceivedAt)
}
