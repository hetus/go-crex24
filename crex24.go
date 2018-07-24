package crex24

import (
	"github.com/hetus/go-crex24/exchange"
)

func New() (e *exchange.Exchange) {
	e = exchange.New()
	return
}
