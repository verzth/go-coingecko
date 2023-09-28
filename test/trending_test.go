package test

import (
	"testing"

	"github.com/verzth/go-coingecko"
)

func TestTrending(t *testing.T) {

	cgClient := gecko.NewClient(nil)

	r, err := cgClient.Trending()
	if r == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
