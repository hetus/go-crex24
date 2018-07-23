package exchange

import (
	"fmt"
)

/*
{
    "symbol": "$PAC",
    "name": "PACCoin",
    "isFiat": false,
    "depositsAllowed": true,
    "depositConfirmationCount": 8,
    "minDeposit": 0.0,
    "withdrawalsAllowed": true,
    "withdrawalPrecision": 8,
    "minWithdrawal": 4.0,
    "maxWithdrawal": null,
    "flatWithdrawalFee": 2.0,
    "isDelisted": false
  }
*/

type Currency struct {
	Symbol                   string      `json:"symbol"`
	Name                     string      `json:"name"`
	IsFiat                   bool        `json:"isFiat"`
	DepositsAllowed          bool        `json:"depositsAllowed"`
	DepositConfirmationCount int64       `json:"depositConfirmationCount"`
	MinDeposit               float64     `json:"minDeposit"`
	WithdrawalsAllowed       bool        `json:"withdrawalsAllowed"`
	WithdrawalPrecision      int64       `json:"withdrawalPrecision"`
	MinWithdrawal            float64     `json:"minWithdrawal"`
	MaxWithdrawal            interface{} `json:"maxWithdrawal"`
	FlatWithdrawalFee        float64     `json:"flatWithdrawalFee"`
	IsDelisted               bool        `json:"isDelisted"`
}

func (c *Currency) String() (s string) {
	s = fmt.Sprintf(
		"%s (%s) Fiat: %t Delisted: %t",
		c.Name, c.Symbol, c.IsFiat, c.IsDelisted,
	)
	return
}

type Currencies []*Currency

func (e *Exchange) Currency(symbol string) (c *Currency, err error) {
	params := EmptyParams()

	var cs Currencies
	err = e.getJSON("/v2/public/currencies?filter="+symbol, params, &cs, false)
	if err == nil && len(cs) > 0 {
		c = cs[0]
	}
	return
}

func (e *Exchange) Currencies() (cs Currencies, err error) {
	params := EmptyParams()
	err = e.getJSON("/v2/public/currencies", params, &cs, false)
	return
}
