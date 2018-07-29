package exchange

import (
	"fmt"
	"strconv"
	"time"
)

// RecentTrade
type RecentTrade struct {
	Price     float64   `json:"price,omitempty"`
	Volume    float64   `json:"volume,omitempty"`
	Side      string    `json:"side,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

// String
func (t *RecentTrade) String() (s string) {
	s = fmt.Sprintf(
		"(Trade) %s = Date: %s Price: %.8f, Volume: %.8f",
		t.Side, t.Timestamp, t.Price, t.Volume,
	)
	return
}

// RecentTrades
type RecentTrades []*RecentTrade

// Trade
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

// Trades
type Trades []*Trade

// TradeFee
type TradeFee struct {
	MakerFeeRate float64   `json:"makerFeeRate,omitempty"`
	TakerFeeRate float64   `json:"takerFeeRate,omitempty"`
	TradeVolume  float64   `json:"tradeVolume,omitempty"`
	LastUpdate   time.Time `json:"lastUpdate,omitempty"`
}

// RecentTrades
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

// TradeFee
func (e *Exchange) TradeFee() (tf TradeFee, err error) {
	params := EmptyParams()
	err = e.postJSON("/v2/trading/tradeFee", params, &tf, true)
	return
}

// TradeHistory
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
