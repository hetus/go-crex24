package exchange

import (
	"fmt"
	"strconv"
	"time"
)

// OBOrder
type OBOrder struct {
	Price  float64 `json:"price"`
	Volume float64 `json:"volume"`
}

// String
func (o *OBOrder) String() (s string) {
	s = fmt.Sprintf(
		"(Order) Price: %.8f, Volume: %.8f",
		o.Price, o.Volume,
	)
	return
}

// OrderBook
type OrderBook struct {
	Buy  []*OBOrder `json:"buyLevels"`
	Sell []*OBOrder `json:"sellLevels"`
}

// OrderTrade
type OrderTrade struct {
	ID          int64   `json:"id,omitempty"`
	OrderID     int64   `json:"orderId,omitempty"`
	Instrument  string  `json:"instrument,omitempty"`
	Side        string  `json:"side,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Volume      float64 `json:"volume,omitempty"`
	Fee         float64 `json:"fee,omitempty"`
	FeeCurrency string  `json:"feeCurrency,omitempty"`
}

// OrderTrades
type OrderTrades []*OrderTrade

// OrderModify
type OrderModify struct {
	ID        int64   `json:"id,omitempty"`
	NewPrice  float64 `json:"newPrice,omitempty"`
	NewVolume float64 `json:"newVolume,omitempty"`
}

// Order
type Order struct {
	ID                 int64       `json:"id,omitempty"`
	Timestamp          time.Time   `json:"timestamp,omitempty"`
	Instrument         string      `json:"instrument,omitempty"`
	Side               string      `json:"side,omitempty"`
	Type               string      `json:"type,omitempty"`
	Status             string      `json:"status,omitempty"`
	CancellationReason string      `json:"cancellationReason,omitempty"`
	TimeInForce        string      `json:"timeInForce,omitempty"`
	Volume             float64     `json:"volume,omitempty"`
	Price              float64     `json:"price,omitempty"`
	StopPrice          interface{} `json:"stopPrice,omitempty"`
	RemainingVolume    float64     `json:"remainingVolume,omitempty"`
	LastUpdate         interface{} `json:"lastUpdate,omitempty"`
	ParentOrderID      interface{} `json:"parentOrderId,omitempty"`
	ChildOrderID       interface{} `json:"childOrderId,omitempty"`
}

// String
func (o *Order) String() (s string) {
	s = fmt.Sprintf(
		"(Order) %d (%s) = Side: %s, Price: %.8f, Volume: %.8f",
		o.ID, o.Instrument, o.Side, o.Price, o.Volume,
	)
	return
}

// Orders
type Orders []*Order

// ActiveOrders
func (e *Exchange) ActiveOrders() (os Orders, err error) {
	params := EmptyParams()
	err = e.getJSON("/v2/trading/activeOrders", params, &os, true)
	return
}

// OrderBook
func (e *Exchange) OrderBook(instrument string, limit int64) (ob OrderBook, err error) {
	if limit < 1 || limit > 1000 {
		limit = 100 // Current API default
	}
	params := EmptyParams()
	err = e.getJSON(
		"/v2/public/orderBook?instrument="+instrument+"&limit="+strconv.FormatInt(limit, 10),
		params, &ob, false,
	)
	return
}

// OrderCancellation
func (e *Exchange) OrderCancellation(id []int64) (ids []int64, err error) {
	err = e.postJSON("/v2/trading/cancelOrdersById", id, ids, true)
	return
}

// OrderHistory
func (e *Exchange) OrderHistory(instrument string, limit int64) (os Orders, err error) {
	if limit < 1 || limit > 1000 {
		limit = 100
	}
	params := EmptyParams()
	err = e.getJSON(
		"/v2/trading/orderHistory?instrument="+instrument+"&limit="+strconv.FormatInt(limit, 10),
		params, &os, true,
	)
	return
}

// OrderModify
func (e *Exchange) OrderModify(om *OrderModify) (o Order, err error) {
	err = e.postJSON("/v2/trading/modifyOrder", om, &o, true)
	return
}

// OrderStatus
func (e *Exchange) OrderStatus(id int64) (o Order, err error) {
	params := EmptyParams()
	err = e.getJSON("/v2/trading/orderStatus?id="+strconv.FormatInt(id, 10), params, &o, true)
	return
}

// OrderTrades
func (e *Exchange) OrderTrades(id int64) (ot OrderTrade, err error) {
	params := EmptyParams()
	err = e.getJSON("/v2/trading/orderTrades?id="+strconv.FormatInt(id, 10), params, &ot, true)
	return
}

// PlaceOrder
func (e *Exchange) PlaceOrder(order *Order) (o Order, err error) {
	err = e.postJSON("/v2/trading/placeOrder", order, &o, true)
	return
}
