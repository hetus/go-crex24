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

type RecentTrade struct {
	Price     float64   `json:"price,omitempty"`
	Volume    float64   `json:"volume,omitempty"`
	Side      string    `json:"side,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (t *RecentTrade) String() (s string) {
	s = fmt.Sprintf(
		"(Trade) %s = Date: %s Price: %.8f, Volume: %.8f",
		t.Side, t.Timestamp, t.Price, t.Volume,
	)
	return
}

type RecentTrades []*RecentTrade

/*
{
    "id": 3005866,
    "orderId": 468533093,
    "timestamp": "2018-06-02T16:26:27Z",
    "instrument": "BCH-ETH",
    "side": "buy",
    "price": 1.78882,
    "volume": 0.027,
    "fee": 0.0000483,
    "feeCurrency": "ETH"
  }
*/

type Trade struct {
	ID          int64     `json:"id"`
	OrderId     int64     `json:"orderId,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
	Instrument  string    `json:"instrument,omitempty"`
	Side        string    `json:"side,omitempty"`
	Price       float64   `json:"price,omitempty"`
	Volume      float64   `json:"volume,omitempty"`
	Fee         float64   `json:"fee,omitempty"`
	FeeCurrency string    `json:"feeCurrency,omitempty"`
}

type Trades []*Trade

func (e *Exchange) RecentTrades(instrument string, limit int64) (ts RecentTrades, err error) {
	if limit < 1 || limit > 1000 {
		limit = 100 // Current API default
	}
	params := EmptyParams()
	err = e.getJSON(
		"/v2/public/recentTrades?instrument="+instrument+"&limit="+strconv.FormatInt(limit, 10),
		params, &ts, false,
	)
	return
}

func (e *Exchange) TradeHistory(instrument string, limit int64) (ts Trades, err error) {
	if limit < 1 || limit > 1000 {
		limit = 100
	}
	params := EmptyParams()
	err = e.getJSON(
		"/v2/trading/tradeHistory?instrument="+instrument+"&limit="+strconv.FormatInt(limit, 10),
		params, &ts, true,
	)
	return
}
