package exchange

import (
	"fmt"
	"strconv"
)

/*
{
    "price": 0.0159019,
    "volume": 0.043717,
    "side": "buy",
    "timestamp": "2018-05-31T10:08:53Z"
  }
*/

type Order struct {
	Price  float64 `json:"price"`
	Volume float64 `json:"volume"`
}

func (t *Order) String() (s string) {
	s = fmt.Sprintf(
		"(Order) Price: %.8f, Volume: %.8f",
		t.Price, t.Volume,
	)
	return
}

type Orders []*Order

type OrderBook struct {
	Buy  []*Order `json:"buyLevels"`
	Sell []*Order `json:"sellLevels"`
}

func (e *Exchange) OrderBook(instrument string, limit int64) (ob OrderBook, err error) {
	if limit < 1 || limit > 1000 {
		limit = 100 // Current API default
	}
	params := EmptyParams()
	err = e.getJSON(
		"/v2/public/orderBook?instrument="+instrument+"&limit="+strconv.FormatInt(limit, 10),
		params, &ob, false,
	)
	return
}
