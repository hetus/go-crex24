package exchange

import (
	"github.com/hetus/go-crex24/client"
	"github.com/hetus/go-crex24/config"
)

// Exchange is an object that provides the API
// for external package usage.
type Exchange struct {
	c   *client.Client
	cfg *config.Config
}

// getJSON is a shortcut method.
func (e *Exchange) getJSON(path string, data, res interface{}, auth bool) (err error) {
	err = e.c.Get(path, data.(map[string]string), res, auth)
	return
}

// postJSON is a shortcut method.
func (e *Exchange) postJSON(path string, data, res interface{}, auth bool) (err error) {
	err = e.c.Post(path, data, res, auth)
	return
}

// New returns a new exchange object that is
// configured according to provide Config object.
func New() (e *Exchange) {
	cfg := config.New()
	e = &Exchange{
		c:   client.New(cfg),
		cfg: cfg,
	}
	return
}
