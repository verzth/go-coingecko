package test

import (
	"testing"

	"github.com/verzth/go-coingecko"
)

func TestSimplePrice(t *testing.T) {
	cgClient := gecko.NewClient(nil)

	r, _ := cgClient.SimplePrice("bitcoin,ethereum", "btc,eth", true, true, true, true)
	if r == nil {
		t.Errorf("Error")
	}

}

func TestSimpleTokenPrice(t *testing.T) {
	cgClient := gecko.NewClient(nil)

	r, _ := cgClient.SimpleTokenPrice("ethereum", "0x0D8775F648430679A709E98d2b0Cb6250d2887EF", "btc,eth", true, true, true, true)
	if r == nil {
		t.Errorf("Error")
	}

}

func TestSimpleSupportedVsCurrency(t *testing.T) {
	cgClient := gecko.NewClient(nil)

	supVcurr, _ := cgClient.SimpleSupportedVsCurrency()
	if supVcurr == nil {
		t.Errorf("Error")
	}
}
