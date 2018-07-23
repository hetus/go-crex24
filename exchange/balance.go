package exchange

import (
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
	Currency  string  `json:"currency"`
	Available float64 `json:"available"`
	Reserved  float64 `json:"reserved"`
}

type Balances []*Balance

func (e *Exchange) Balance(currency string) (c *Balance, err error) {
	params := EmptyParams()

	var bs Balances
	err = e.getJSON("/public/balances?currency="+currency+"&nonZeroOnly=false", params, &bs, true)
	if err == nil && len(bs) > 0 {
		c = bs[0]
	}
	return
}

func (e *Exchange) Balances(nonZeroOnly bool) (bs Balances, err error) {
	params := EmptyParams()
	err = e.getJSON("/public/balances?nonZeroOnly="+strconv.FormatBool(nonZeroOnly), params, &bs, true)
	return
}
