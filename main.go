package main

import (
	"github.com/hetus/go-crex24/crex24"
)

func New() (c24 *crex24.Exchange) {
	c24 = crex24.New()
	return
}
