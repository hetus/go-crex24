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

/*
{
  "warning": null,
  "balanceDeduction": 1.5005,
  "fee": 0.0005,
  "payout": 1.5
}
*/

type WithdrawalPreview struct {
	Warning          interface{} `json:"warning,omitempty"`
	BalanceDeduction float64     `json:"balanceDeduction,omitempty"`
	Fee              float64     `json:"fee,omitempty"`
	Payout           float64     `json:"payout,omitempty"`
}

/*
{
  "id": 737551,
  "type": "withdrawal",
  "currency": "ETH",
  "address": "0x184189a9187c918ef91875641f9781a9187b75a7",
  "paymentId": "",
  "amount": 54.1,
  "fee": 0.005,
  "txId": "0x1983645416f16a16c1687643086f7c91767a9817b981765140c8176871fc79fa",
  "createdAt": "2018-06-01T06:48:32Z",
  "processedAt": "2018-06-01T07:20:14Z",
  "confirmationsRequired": 12,
  "confirmationCount": 12,
  "status": "success",
  "errorDescription": null
}
*/

type Withdrawal struct {
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

type WithdrawalRequest struct {
	Currency   string      `json:"currency,omitempty"`
	Address    string      `json:"address,omitempty"`
	PaymentID  interface{} `json:"paymentId,omitempty"`
	Amount     float64     `json:"amount,omitempty"`
	IncludeFee bool        `json:"includeFee,omitempty"`
}

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

func (e *Exchange) MoneyTransferStatus(ids string) (ms *MoneyTransfers, err error) {
	params := EmptyParams()
	err = e.getJSON("/v2/account/moneyTransferStatus?id="+ids, params, &ms, true)
	return
}

func (e *Exchange) Withdrawal(wr *WithdrawalRequest) (mt *MoneyTransfer, err error) {
	err = e.postJSON("/v2/account/withdraw", wr, &mt, true)
	return
}

func (e *Exchange) WithdrawalPreview(currency string, amount float64, includeFee bool) (wp *WithdrawalPreview, err error) {
	params := EmptyParams()
	err = e.getJSON(
		"/v2/account/previewWithdrawal?currency="+currency+"&amount="+fmt.Sprintf("%.f", amount)+"&includeFee="+strconv.FormatBool(includeFee),
		params, &wp, includeFee,
	)
	return
}
