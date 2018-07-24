# go-crex24
Go API interface to Crex24 exchange API v2.

__NOTE:__ Under active development and subject to change.


## Install
`go get -u github.com/hetus/go-crex24`


## Configure
To configure you need to copy `.env.example` to `.env` in your local folder and change the API `KEY` and `SECRET` fields.


## Run Demo
`go run cmd/crex24-demo/main.go`


## Use
Example usage:
```go
package main

import (
	"fmt"
	"log"

	"github.com/hetus/go-crex24"
)

func main() {
    api := crex24.New()

    ts, err := api.Tickers()
    if err != nil {
        log.Fatalf("Error: %v\n", err)
    }
    fmt.Println(ts)
}
```


## TODO
- Finished adding v2 endpoints
- Added validation to endpoints
- Write tests for endpoints
- Release to public
- Contact Crex24