package coins

import "github.com/verzth/go-coingecko/types"

type Tickers struct {
	Name    string         `json:"name"`
	Tickers []types.Ticker `json:"tickers"`
}
