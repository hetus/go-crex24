package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	httpclient "github.com/ddliu/go-httpclient"
	"github.com/methuselahdev/go-crex24/auth"
	"github.com/methuselahdev/go-crex24/config"
)

// Response for requests.
type Response httpclient.Response

// Client is the http client wrapper for Crex24.
type Client struct {
	api     *httpclient.HttpClient
	debug   bool
	lastAt  time.Time
	nonce   int64
	secret  string
	url     string
	version string
	wait    time.Duration
}

// Get will make a GET request with provided path and params.
func (c *Client) Get(path string, params map[string]string, auth bool) (res Response, err error) {
	if c.debug {
		fmt.Println("GET:", auth, path, params)
	}
	if auth {
		c.Sign(path, "")
	}

	var r *httpclient.Response
	r, err = c.api.Get(c.URL()+path, params)
	if err != nil {
		return
	}

	res = Response(*r)
	return
}

// Nonce increments the nonce and sets the header.
func (c *Client) Nonce() (n string) {
	c.nonce++
	n = strconv.FormatInt(c.nonce, 10)
	c.api.Headers["X-CREX24-API-NONCE"] = n
	if c.debug {
		fmt.Println("X-CREX24-API-NONCE:", n)
	}
	return
}

// Post will make a POST request.
func (c *Client) Post(path string, data interface{}, auth bool) (res Response, err error) {
	if c.debug {
		fmt.Println("POST:", auth, path, data)
	}
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(data)
	if err != nil {
		return
	}
	if auth {
		c.Sign(path, buf.String())
	}

	var r *httpclient.Response
	r, err = c.api.PostJson(c.URL()+path, data)
	if err != nil {
		return
	}

	res = Response(*r)
	return
}

// Sign adds the authentication information to the header.
func (c *Client) Sign(path, data string) (err error) {
	nonce := c.Nonce()
	if c.debug {
		fmt.Println("Sign:", path, nonce, data)
	}

	var msg bytes.Buffer
	_, err = msg.WriteString(path)
	_, err = msg.WriteString(nonce)
	if data != "" {
		_, err = msg.WriteString(data)
	}
	if err != nil {
		return
	}

	var signMsg string
	signMsg, err = auth.Sign(msg.Bytes(), c.secret)
	if err != nil {
		return
	}
	c.api.Headers["X-CREX24-API-SIGN"] = signMsg
	if c.debug {
		fmt.Println("X-CREX24-API-SIGN:", signMsg)
	}
	return
}

// URL returns the api host plus version.
func (c *Client) URL() (url string) {
	url = c.url + "/" + c.version
	return
}

// New creates a new client and returns a pointer.
func New(cfg *config.Config) (c *Client) {
	c = &Client{
		api:     httpclient.NewHttpClient(),
		debug:   cfg.Debug,
		lastAt:  time.Now().UTC(),
		secret:  cfg.APISecret,
		url:     cfg.APIUrl,
		version: cfg.APIVersion,
		wait:    1 * time.Second,
	}

	c.api.Defaults(httpclient.Map{
		"Accept":                 "application/json",
		"Content-Type":           "application/json",
		httpclient.OPT_USERAGENT: "go-crex24",
		"X-CREX24-API-KEY":       cfg.APIKey,
	})
	if c.debug {
		fmt.Println("X-CREX24-API-KEY:", c.api.Headers["X-CREX24-API-KEY"])
		fmt.Println("Client.api.Headers:", c.api.Headers)
	}
	return
}
