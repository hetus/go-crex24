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
	Symbol                   string      `json:"symbol,omitempty"`
	Name                     string      `json:"name,omitempty"`
	IsFiat                   bool        `json:"isFiat,omitempty"`
	DepositsAllowed          bool        `json:"depositsAllowed,omitempty"`
	DepositConfirmationCount int64       `json:"depositConfirmationCount,omitempty"`
	MinDeposit               float64     `json:"minDeposit,omitempty"`
	WithdrawalsAllowed       bool        `json:"withdrawalsAllowed,omitempty"`
	WithdrawalPrecision      int64       `json:"withdrawalPrecision,omitempty"`
	MinWithdrawal            float64     `json:"minWithdrawal,omitempty"`
	MaxWithdrawal            interface{} `json:"maxWithdrawal,omitempty"`
	FlatWithdrawalFee        float64     `json:"flatWithdrawalFee,omitempty"`
	IsDelisted               bool        `json:"isDelisted,omitempty"`
}

func (c *Currency) String() (s string) {
	s = fmt.Sprintf(
		"(Currency) %s (%s) = Fiat: %t, Delisted: %t",
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
