package main

import (
	"fmt"

	"github.com/hetus/go-crex24/exchange"
)

func main() {
	e := exchange.New()

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
	fmt.Printf("BTC: %v\n\n", c)

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
	fmt.Printf("LTC-BTC: %v\n\n", i)

	i, err = e.Instrument("BTC-LTC")
	if err != nil {
		fmt.Printf("Instrument (BTC-LTC): %v\n", err)
	}
	fmt.Printf("BTC-LTC: %v\n\n", i)

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
	fmt.Printf("BTC Balance: %v\n\n", b)
}
