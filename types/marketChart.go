package types

import "github.com/shopspring/decimal"

type MarketChart struct {
	Prices       [][]decimal.Decimal `json:"prices"`
	MarketCaps   [][]decimal.Decimal `json:"market_caps"`
	TotalVolumes [][]decimal.Decimal `json:"total_volumes"`
}
