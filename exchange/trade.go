package exchange

import (
	"fmt"
	"strconv"
	"time"
)

/*
{
    "price": 0.0159019,
    "volume": 0.043717,
    "side": "buy",
    "timestamp": "2018-05-31T10:08:53Z"
  }
*/

type Trade struct {
	Price     float64   `json:"price"`
	Volume    float64   `json:"volume"`
	Side      string    `json:"side"`
	Timestamp time.Time `json:"timestamp"`
}

func (t *Trade) String() (s string) {
	s = fmt.Sprintf(
		"(Trade) %s = Date: %s Price: %.8f, Volume: %.8f",
		t.Side, t.Timestamp, t.Price, t.Volume,
	)
	return
}

type Trades []*Trade

func (e *Exchange) RecentTrades(instrument string, limit int64) (is Trades, err error) {
	if limit < 1 || limit > 1000 {
		limit = 100 // Current API default
	}
	params := EmptyParams()
	err = e.getJSON(
		"/v2/public/recentTrades?instrument="+instrument+"&limit="+strconv.FormatInt(limit, 10),
		params, &is, false,
	)
	return
}
