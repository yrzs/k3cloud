package server

import (
	"github.com/yrzs/k3cloud/kernel"
)

// NewK3Cloud. new application
func NewK3Cloud(c *kernel.K3Config) (*kernel.K3Cloud, error) {
	browser := kernel.NewBrowserWithTransport()
	//err := browser.InitLogin(c)
	//if err != nil {
	//	return nil, err
	//}
	app := &kernel.K3Cloud{
		Config: c,
		Client: browser,
	}
	return app, nil
}
