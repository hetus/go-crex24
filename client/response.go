package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"

	httpclient "github.com/ddliu/go-httpclient"
)

// ErrorResponse for errors from API.
type ErrorResponse struct {
	Message string `json:"errorDescription"`
}

// handleResponse is a helper function to:
//	- trim BOM from utf8 prefix
// 	- change error response to error message
func handleResponse(req *httpclient.Response, res interface{}, auth bool) (err error) {
	var b []byte
	b, err = ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}

	if auth {
		b = bytes.TrimPrefix(b, []byte("\xef\xbb\xbf"))
	}

	err = json.NewDecoder(bytes.NewReader(b)).Decode(res)
	if err != nil {
		var res2 ErrorResponse
		err2 := json.NewDecoder(bytes.NewReader(b)).Decode(&res2)
		if err2 == nil {
			err = errors.New(res2.Message)
		}
	}
	return
}
