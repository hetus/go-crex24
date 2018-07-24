package main

import (
	"fmt"

	"github.com/hetus/go-crex24/exchange"
)

func main() {
	e := exchange.New()

	/** Public */
	fmt.Printf("\t*** PUBLIC ***\n\n")

	// Currency
	cs, err := e.Currencies()
	if err != nil {
		fmt.Printf("(Error) Currencies: %v\n", err)
	}
	fmt.Printf("Currencies: %d\n\n", len(cs))

	c, err := e.Currency("BTC")
	if err != nil {
		fmt.Printf("(Error) Currency (BTC): %v\n", err)
	}
	fmt.Printf("%v\n\n", c)

	// Instrument
	is, err := e.Instruments()
	if err != nil {
		fmt.Printf("(Error) Instruments: %v\n", err)
	}
	fmt.Printf("Instruments: %d\n\n", len(is))

	i, err := e.Instrument("LTC-BTC")
	if err != nil {
		fmt.Printf("(Error) Instrument (LTC-BTC): %v\n", err)
	}
	fmt.Printf("%v\n\n", i)

	i, err = e.Instrument("BTC-LTC")
	if err != nil {
		fmt.Printf("(Error) Instrument (BTC-LTC): %v\n\n", err)
	} else {
		fmt.Printf("%v\n\n", i)
	}

	// Ticker
	ts, err := e.Tickers()
	if err != nil {
		fmt.Printf("(Error) Tickers: %v\n", err)
	}
	fmt.Printf("Tickers: %d\n\n", len(ts))

	t, err := e.Ticker("BTC-USD")
	if err != nil {
		fmt.Printf("(Error) Ticker (BTC-USD): %v\n", err)
	}
	fmt.Printf("%v\n\n", t)

	// Trade
	tss, err := e.RecentTrades("LTC-BTC", 5)
	if err != nil {
		fmt.Printf("(Error) Recent Trades (LTC-BTC): %v\n", err)
	}
	fmt.Printf("Recent Trades (LTC-BTC): %d\n", len(tss))
	fmt.Printf("%v\n\n", tss[0])

	// Order
	ob, err := e.OrderBook("LTC-BTC", 5)
	if err != nil {
		fmt.Printf("(Error) Order Book (LTC-BTC): %v\n", err)
	}
	fmt.Printf("Order Book (LTC-BTC): Buys: %d, Sells: %d\n", len(ob.Buy), len(ob.Sell))
	fmt.Printf("Buy: %v\nSell: %v\n\n", ob.Buy[0], ob.Sell[0])

	/** Auth Required */
	fmt.Printf("\t*** AUTH REQUIRED ***\n\n")

	// Order
	o := exchange.Order{
		Instrument: "LTC-BTC",
		Side:       "buy",
		Price:      0.00000001,
		Volume:     1.0,
	}
	no, err := e.PlaceOrder(&o)
	if err != nil {
		fmt.Printf("(Error) Place Order (LTC-BTC): %v\n\n", err)
	} else {
		fmt.Printf("%v\n\n", &no)
	}

	os, err := e.OrderStatus(o.ID)
	if err != nil {
		fmt.Printf("(Error) Order Status: %v\n\n", err)
	} else {
		fmt.Printf("%v\n\n", &os)
	}

	// Balance
	bs, err := e.Balances(true)
	if err != nil {
		fmt.Printf("(Error) Balances: %v\n", err)
	}
	fmt.Printf("Balances: %d\n\n", len(bs))

	b, err := e.Balance("BTC")
	if err != nil {
		fmt.Printf("(Error) Balance (BTC): %v\n", err)
	}
	fmt.Printf("%v\n\n", b)
}
