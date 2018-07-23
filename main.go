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
	fmt.Println("Currencies:", len(cs))

	c, err := e.Currency("BTC")
	if err != nil {
		log.Fatalf("Currency (BTC): %v", err)
	}
	fmt.Println("BTC:", c)

	// Balance
	bs, err := e.Balances(false)
	if err != nil {
		log.Fatalf("Balances: %v", err)
	}
	fmt.Println("Balances:", len(bs))

	b, err := e.Balance("BTC")
	if err != nil {
		log.Fatalf("Balance (BTC): %v", err)
	}
	fmt.Println("BTC Balance:", b)
}
