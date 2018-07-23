package main

import (
	"fmt"
	"log"

	"github.com/hetus/go-crex24/exchange"
)

func main() {
	e := exchange.New()

	// Currency
	cs, err := e.Currencies()
	if err != nil {
		log.Fatalf("Currencies: %v", err)
	}
	fmt.Println("Currencies:", len(cs), "\n")

	c, err := e.Currency("BTC")
	if err != nil {
		log.Fatalf("Currency (BTC): %v", err)
	}
	fmt.Println("BTC:", c, "\n")

	// Instrument
	is, err := e.Instruments()
	if err != nil {
		log.Fatalf("Instruments: %v", err)
	}
	fmt.Println("Instruments:", len(is), "\n")

	i, err := e.Instrument("LTC-BTC")
	if err != nil {
		log.Fatalf("Instrument (LTC-BTC): %v", err)
	}
	fmt.Println("LTC-BTC:", i, "\n")

	// Balance
	bs, err := e.Balances(true)
	if err != nil {
		log.Fatalf("Balances: %v", err)
	}
	fmt.Println("Balances:", len(bs), "\n")

	b, err := e.Balance("BTC")
	if err != nil {
		log.Fatalf("Balance (BTC): %v", err)
	}
	fmt.Println("BTC Balance:", b, "\n")
}
