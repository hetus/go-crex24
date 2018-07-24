package crex24

import (
	"fmt"
	"strconv"
)

/*
{
    "currency": "ALPS",
    "available": 0.0979,
    "reserved": 0.0
  }
*/

type Balance struct {
	Currency  string  `json:"currency,omitempty"`
	Available float64 `json:"available,omitempty"`
	Reserved  float64 `json:"reserved,omitempty"`
}

func (b *Balance) String() (s string) {
	s = fmt.Sprintf(
		"(Balance) %s = Available: %.8f, Reserved: %.8f",
		b.Currency, b.Available, b.Reserved,
	)
	return
}

type Balances []*Balance

func (e *Exchange) Balance(currency string) (b *Balance, err error) {
	params := EmptyParams()

	var bs Balances
	err = e.getJSON("/v2/account/balance?currency="+currency+"&nonZeroOnly=false", params, &bs, true)
	if err == nil && len(bs) > 0 {
		b = bs[0]
	}
	return
}

func (e *Exchange) Balances(nonZeroOnly bool) (bs Balances, err error) {
	params := EmptyParams()
	err = e.getJSON("/v2/account/balance?nonZeroOnly="+strconv.FormatBool(nonZeroOnly), params, &bs, true)
	return
}
