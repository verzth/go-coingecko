package test

import (
	"fmt"
	"testing"

	"github.com/verzth/go-coingecko"
)

func TestCoins(t *testing.T) {
	coin := "bitcoin"

	cgClient := gecko.NewClient(nil)

	coinData, _ := cgClient.CoinsId(coin, true, true, true, true, true, true)
	fmt.Println(coinData.MarketData.CurrentPrice.Usd)

}
