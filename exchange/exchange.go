package exchange

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/hetus/go-crex24/client"
	"github.com/hetus/go-crex24/config"
)

type ErrorResponse struct {
	Message string `json:"errorDescription"`
}

type Exchange struct {
	c   *client.Client
	cfg *config.Config
}

func (e *Exchange) getJSON(path string, data, res interface{}, auth bool) (err error) {
	var r client.Response
	r, err = e.c.Get(path, data.(map[string]string), auth)
	if err != nil {
		return
	}
	defer r.Body.Close()

	var b []byte
	b, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	if auth {
		b = bytes.TrimPrefix(b, []byte("\xef\xbb\xbf"))
	}

	decoder := json.NewDecoder(bytes.NewReader(b))
	err = decoder.Decode(res)
	if err != nil {
		var res2 ErrorResponse
		err2 := decoder.Decode(&res2)
		if err2 == nil {
			err = errors.New(res2.Message)
		}
	}
	return
}

func (e *Exchange) postJSON(path string, data, res interface{}, auth bool) (err error) {
	var r client.Response
	r, err = e.c.Post(path, data, auth)
	if err != nil {
		return
	}
	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(res)
	return
}

func New() (e *Exchange) {
	cfg := config.New()
	e = &Exchange{
		c:   client.New(cfg),
		cfg: cfg,
	}
	return
}
