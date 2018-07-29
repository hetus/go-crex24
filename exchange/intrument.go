package exchange

import "fmt"

// Instrument
type Instrument struct {
	Symbol              string   `json:"symbol,omitempty"`
	BaseCurrency        string   `json:"baseCurrency,omitempty"`
	QuoteCurrency       string   `json:"quoteCurrency,omitempty"`
	FeeCurrency         string   `json:"feeCurrency,omitempty"`
	TickSize            float64  `json:"tickSize,omitempty"`
	MinPrice            float64  `json:"minPrice,omitempty"`
	MinVolume           float64  `json:"minVolume,omitempty"`
	SupportedOrderTypes []string `json:"supportedOrderTypes,omitempty"`
	State               string   `json:"state,omitempty"`
}

// String
func (i *Instrument) String() (s string) {
	s = fmt.Sprintf(
		"(Instrument) %s = Base: %s, Quote: %s, Tick Size: %.8f, Min. Price: %.8f, Min. Volume: %.8f",
		i.Symbol, i.BaseCurrency, i.QuoteCurrency, i.TickSize, i.MinPrice, i.MinVolume,
	)
	return
}

// Instruments
type Instruments []*Instrument

// Instrument
func (e *Exchange) Instrument(symbol string) (i *Instrument, err error) {
	params := EmptyParams()

	var is Instruments
	err = e.getJSON("/v2/public/instruments?filter="+symbol, params, &is, false)
	if err == nil && len(is) > 0 {
		i = is[0]
	}
	return
}

// Instruments
func (e *Exchange) Instruments() (is Instruments, err error) {
	params := EmptyParams()
	err = e.getJSON("/v2/public/instruments", params, &is, false)
	return
}
