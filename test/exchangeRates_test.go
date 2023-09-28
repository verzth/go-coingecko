package test

import (
	"testing"

	"github.com/verzth/go-coingecko"
)

func TestExchangeRates(t *testing.T) {

	cgClient := gecko.NewClient(nil)

	r, err := cgClient.ExchangeRates()
	if r == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
