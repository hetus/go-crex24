package exchange

/*
{
    "symbol": "$PAC-BTC",
    "baseCurrency": "$PAC",
    "quoteCurrency": "BTC",
    "feeCurrency": "BTC",
    "tickSize": 0.00000001,
    "minPrice": 0.00000001,
    "minVolume": 1.0,
    "supportedOrderTypes": [
      "limit"
    ],
    "state": "active"
  }
*/

type Instrument struct {
	Symbol              string   `json:"symbol"`
	BaseCurrency        string   `json:"baseCurrency"`
	QuoteCurrency       string   `json:"quoteCurrency"`
	FeeCurrency         string   `json:"feeCurrency"`
	TickSize            float64  `json:"tickSize"`
	MinPrice            float64  `json:"minPrice"`
	MinVolume           float64  `json:"minVolume"`
	SupportedOrderTypes []string `json:"supportedOrderTypes"`
	State               string   `json:"state"`
}

type Instruments []*Instrument

func (e *Exchange) Instrument(symbol string) (i *Instrument, err error) {
	params := EmptyParams()

	var is Instruments
	err = e.getJSON("/v2/public/instruments?filter="+symbol, params, &is, false)
	if err == nil && len(is) > 0 {
		i = is[0]
	}
	return
}

func (e *Exchange) Instruments() (is Instruments, err error) {
	params := EmptyParams()
	err = e.getJSON("/v2/public/instruments", params, &is, false)
	return
}
