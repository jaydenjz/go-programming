package hello

import (
	"rsc.io/quote/v3"
)

// Hello returns quote hello
func Hello() string {
	return quote.HelloV3()
}

// Proverb returns quote Concurrency
func Proverb() string {
	return quote.Concurrency()
}
