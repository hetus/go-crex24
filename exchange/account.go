package exchange

import (
	"fmt"
	"strconv"
	"time"
)

// Balance
type Balance struct {
	Currency  string  `json:"currency,omitempty"`
	Available float64 `json:"available,omitempty"`
	Reserved  float64 `json:"reserved,omitempty"`
}

// String
func (b *Balance) String() (s string) {
	s = fmt.Sprintf(
		"(Balance) %s = Available: %.8f, Reserved: %.8f",
		b.Currency, b.Available, b.Reserved,
	)
	return
}

// Balances
type Balances []*Balance

// DepositAddress
type DepositAddress struct {
	Currency  string      `json:"currency,omitempty"`
	Address   string      `json:"address,omitempty"`
	PaymentId interface{} `json:"paymentId,omitempty"`
}

// MoneyTransfer
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

// MoneyTransfers
type MoneyTransfers []*MoneyTransfer

// WithdrawalPreview
type WithdrawalPreview struct {
	Warning          interface{} `json:"warning,omitempty"`
	BalanceDeduction float64     `json:"balanceDeduction,omitempty"`
	Fee              float64     `json:"fee,omitempty"`
	Payout           float64     `json:"payout,omitempty"`
}

// Withdrawal
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

// WithdrawalRequest
type WithdrawalRequest struct {
	Currency   string      `json:"currency,omitempty"`
	Address    string      `json:"address,omitempty"`
	PaymentID  interface{} `json:"paymentId,omitempty"`
	Amount     float64     `json:"amount,omitempty"`
	IncludeFee bool        `json:"includeFee,omitempty"`
}

// Balance
func (e *Exchange) Balance(currency string) (b *Balance, err error) {
	params := EmptyParams()

	var bs Balances
	err = e.getJSON("/v2/account/balance?currency="+currency+"&nonZeroOnly=false", params, &bs, true)
	if err == nil && len(bs) > 0 {
		b = bs[0]
	}
	return
}

// Balances
func (e *Exchange) Balances(nonZeroOnly bool) (bs Balances, err error) {
	params := EmptyParams()
	err = e.getJSON("/v2/account/balance?nonZeroOnly="+strconv.FormatBool(nonZeroOnly), params, &bs, true)
	return
}

// DepositAddress
func (e *Exchange) DepositAddress(currency string) (da *DepositAddress, err error) {
	params := EmptyParams()
	err = e.getJSON("/v2/account/depositAddress?currency="+currency, params, &da, true)
	return
}

// MoneyTransfers
func (e *Exchange) MoneyTransfers(currency, t string) (ms *MoneyTransfers, err error) {
	params := EmptyParams()
	err = e.getJSON("/v2/account/moneyTransfers?currency="+currency+"&type="+t, params, &ms, true)
	return
}

// MoneyTransferStatus
func (e *Exchange) MoneyTransferStatus(ids string) (ms *MoneyTransfers, err error) {
	params := EmptyParams()
	err = e.getJSON("/v2/account/moneyTransferStatus?id="+ids, params, &ms, true)
	return
}

// Withdrawal
func (e *Exchange) Withdrawal(wr *WithdrawalRequest) (mt *MoneyTransfer, err error) {
	err = e.postJSON("/v2/account/withdraw", wr, &mt, true)
	return
}

// WithdrawalPreview
func (e *Exchange) WithdrawalPreview(currency string, amount float64, includeFee bool) (wp *WithdrawalPreview, err error) {
	params := EmptyParams()
	err = e.getJSON(
		"/v2/account/previewWithdrawal?currency="+currency+"&amount="+fmt.Sprintf("%.f", amount)+"&includeFee="+strconv.FormatBool(includeFee),
		params, &wp, includeFee,
	)
	return
}
