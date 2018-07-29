package exchange

import (
	"fmt"
	"time"
)

// Ticker
type Ticker struct {
	Instrument    string    `json:"instrument,omitempty"`
	Last          float64   `json:"last,omitempty"`
	PercentChange float64   `json:"percentChange,omitempty"`
	Low           float64   `json:"low,omitempty"`
	High          float64   `json:"high,omitempty"`
	BaseVolume    float64   `json:"baseVolume,omitempty"`
	QuoteVolume   float64   `json:"quoteVolume,omitempty"`
	VolumeInBtc   float64   `json:"volumeInBtc,omitempty"`
	VolumeInUsd   float64   `json:"volumeInUsd,omitempty"`
	Ask           float64   `json:"ask,omitempty"`
	Bid           float64   `json:"bid,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

// String
func (t *Ticker) String() (s string) {
	s = fmt.Sprintf(
		"(Ticker) %s = Last: %.8f, High: %.8f, Low: %.8f, Ask: %.8f, Bid: %.8f, Volume (BTC): %.8f",
		t.Instrument, t.Last, t.High, t.Low, t.Ask, t.Bid, t.VolumeInBtc,
	)
	return
}

// Tickers
type Tickers []*Ticker

// Ticker
func (e *Exchange) Ticker(instrument string) (t *Ticker, err error) {
	params := EmptyParams()

	var ts Tickers
	err = e.getJSON("/v2/public/tickers?instrument="+instrument, params, &ts, false)
	if err == nil && len(ts) > 0 {
		t = ts[0]
	}
	return
}

// Tickers
func (e *Exchange) Tickers() (ts Tickers, err error) {
	params := EmptyParams()
	err = e.getJSON("/v2/public/tickers", params, &ts, false)
	return
}
