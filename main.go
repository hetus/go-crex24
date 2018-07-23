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
		fmt.Printf("Currencies: %v", err)
	}
	fmt.Printf("Currencies: %d\n\n", len(cs))

	c, err := e.Currency("BTC")
	if err != nil {
		fmt.Printf("Currency (BTC): %v\n", err)
	}
	fmt.Printf("%v\n\n", c)

	// Instrument
	is, err := e.Instruments()
	if err != nil {
		fmt.Printf("Instruments: %v\n", err)
	}
	fmt.Printf("Instruments: %d\n\n", len(is))

	i, err := e.Instrument("LTC-BTC")
	if err != nil {
		fmt.Printf("Instrument (LTC-BTC): %v\n", err)
	}
	fmt.Printf("%v\n\n", i)

	i, err = e.Instrument("BTC-LTC")
	if err != nil {
		fmt.Printf("(Error) Instrument (BTC-LTC): %v\n", err)
	} else {
		fmt.Printf("%v\n\n", i)
	}

	// Ticker
	ts, err := e.Tickers()
	if err != nil {
		fmt.Printf("Tickers: %v", err)
	}
	fmt.Printf("Tickers: %d\n\n", len(ts))

	t, err := e.Ticker("BTC-USD")
	if err != nil {
		fmt.Printf("Ticker (BTC-USD): %v\n", err)
	}
	fmt.Printf("%v\n\n", t)

	/** Auth Required */
	fmt.Printf("\t*** AUTH REQUIRED ***\n\n")

	// Balance
	bs, err := e.Balances(true)
	if err != nil {
		fmt.Printf("Balances: %v\n", err)
	}
	fmt.Printf("Balances: %d\n\n", len(bs))

	b, err := e.Balance("BTC")
	if err != nil {
		fmt.Printf("Balance (BTC): %v\n", err)
	}
	fmt.Printf("%v\n\n", b)
}
