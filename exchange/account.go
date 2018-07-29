package exchange

import (
	"fmt"
	"strconv"
	"time"
)

/*
{
    "currency": "ALPS",
    "available": 0.0979,
    "reserved": 0.0
  }
*/

type Balance struct {
	Currency  string  `json:"currency,omitempty"`
	Available float64 `json:"available,omitempty"`
	Reserved  float64 `json:"reserved,omitempty"`
}

func (b *Balance) String() (s string) {
	s = fmt.Sprintf(
		"(Balance) %s = Available: %.8f, Reserved: %.8f",
		b.Currency, b.Available, b.Reserved,
	)
	return
}

type Balances []*Balance

/*
{
  "currency": "BTC",
  "address": "5xF3EVqwOf53PLeU78iGJpbWz45qzPIfnd",
  "paymentId": null
}
*/

type DepositAddress struct {
	Currency  string      `json:"currency,omitempty"`
	Address   string      `json:"address,omitempty"`
	PaymentId interface{} `json:"paymentId,omitempty"`
}

/*
{
    "id": 756446,
    "type": "deposit",
    "currency": "ETH",
    "address": "0x451d5a1b7519aa75164f440df78c74aac96023fe",
    "paymentId": null,
    "amount": 0.142,
    "fee": null,
    "txId": "0x2b49098749840a9482c4894be94f94864b498a1306b6874687a5640cc9871918",
    "createdAt": "2018-06-02T19:30:28Z",
    "processedAt": "2018-06-02T21:10:41Z",
    "confirmationsRequired": 12,
    "confirmationCount": 12,
    "status": "success",
    "errorDescription": null
  }
*/

type MoneyTransfer struct {
	ID                    int64       `json:"id,omitempty"`
	Type                  string      `json:"type,omitempty"`
	Currency              string      `json:"currency,omitempty"`
	Address               string      `json:"address,omitempty"`
	PaymentID             interface{} `json:"paymentId,omitempty"`
	Amount                float64     `json:"amount,omitempty"`
	Fee                   interface{} `json:"fee,omitempty"`
	TxID                  string      `json:"txId,omitempty"`
	CreatedAt             time.Time   `json:"createdAt,omitempty"`
	ProcessedAt           time.Time   `json:"processedAt,omitempty"`
	ConfirmationsRequired int64       `json:"confirmationsRequired,omitempty"`
	ConfirmationCount     int64       `json:"confirmationCount,omitempty"`
	Status                string      `json:"status,omitempty"`
	ErrorDescription      interface{} `json:"errorDescription,omitempty"`
}

type MoneyTransfers []*MoneyTransfer

func (e *Exchange) Balance(currency string) (b *Balance, err error) {
	params := EmptyParams()

	var bs Balances
	err = e.getJSON("/v2/account/balance?currency="+currency+"&nonZeroOnly=false", params, &bs, true)
	if err == nil && len(bs) > 0 {
		b = bs[0]
	}
	return
}

func (e *Exchange) Balances(nonZeroOnly bool) (bs Balances, err error) {
	params := EmptyParams()
	err = e.getJSON("/v2/account/balance?nonZeroOnly="+strconv.FormatBool(nonZeroOnly), params, &bs, true)
	return
}

func (e *Exchange) DepositAddress(currency string) (da *DepositAddress, err error) {
	params := EmptyParams()
	err = e.getJSON("/v2/account/depositAddress?currency="+currency, params, &da, true)
	return
}

func (e *Exchange) MoneyTransfers(currency, t string) (ms *MoneyTransfers, err error) {
	params := EmptyParams()
	err = e.getJSON("/v2/account/moneyTransfers?currency="+currency+"&type="+t, params, &ms, true)
	return
}
