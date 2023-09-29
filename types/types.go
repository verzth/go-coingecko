package types

import "github.com/shopspring/decimal"

type Description struct {
	En   string `json:"en"`
	De   string `json:"de"`
	Es   string `json:"es"`
	Fr   string `json:"fr"`
	It   string `json:"it"`
	Pl   string `json:"pl"`
	Ro   string `json:"ro"`
	Hu   string `json:"hu"`
	Nl   string `json:"nl"`
	Pt   string `json:"pt"`
	Sv   string `json:"sv"`
	Vi   string `json:"vi"`
	Tr   string `json:"tr"`
	Ru   string `json:"ru"`
	Ja   string `json:"ja"`
	Zh   string `json:"zh"`
	Ko   string `json:"ko"`
	Ar   string `json:"ar"`
	Th   string `json:"th"`
	ID   string `json:"id"`
	Zhtw string `json:"zh-tw"`
}

type Links struct {
	Homepage                    []string        `json:"homepage"`
	BlockchainSite              []string        `json:"blockchain_site"`
	OfficialForumURL            []string        `json:"official_forum_url"`
	ChatURL                     []string        `json:"chat_url"`
	AnnouncementURL             []string        `json:"announcement_url"`
	TwitterScreenName           string          `json:"twitter_screen_name"`
	FacebookUsername            string          `json:"facebook_username"`
	BitcointalkThreadIdentifier decimal.Decimal `json:"bitcointalk_thread_identifier"`
	TelegramChannelIdentifier   string          `json:"telegram_channel_identifier"`
	SubredditURL                string          `json:"subreddit_url"`
	ReposURL                    ReposUrl        `json:"repos_url"`
}

type ReposUrl struct {
	Github    []string `json:"github"`
	Bitbucket []string `json:"bitbucket"`
}

type Roi struct {
	Time       decimal.Decimal `json:"times"`
	Currency   string          `json:"currency"`
	Percentage decimal.Decimal `json:"percentage"`
}

type Sparkline struct {
	Price []decimal.Decimal `json:"price"`
}

type Ticker struct {
	Base                   string          `json:"base"`
	Target                 string          `json:"target"`
	Market                 MarketTicker    `json:"market"`
	Last                   decimal.Decimal `json:"last"`
	Volume                 decimal.Decimal `json:"volume"`
	ConvertedLast          Converted       `json:"converted_last"`
	ConvertedVolume        Converted       `json:"converted_volume"`
	TrustScore             string          `json:"trust_score"`
	BidAskSpreadPercentage decimal.Decimal `json:"bid_ask_spread_percentage"`
	Timestamp              string          `json:"timestamp"`
	LastTradedAt           string          `json:"last_traded_at"`
	LastFetchAt            string          `json:"last_fetch_at"`
	IsAnomaly              bool            `json:"is_anomaly"`
	IsStale                bool            `json:"is_stale"`
	TradeURL               string          `json:"trade_url"`
	// token_info_url `json:"token_info_url"`	null
	CoinID string `json:"coin_id"`
}

type MarketTicker struct {
	Name                string `json:"name"`
	Identifier          string `json:"identifier"`
	HasTradingIncentive bool   `json:"has_trading_incentive"`
}

type Converted struct {
	Btc decimal.Decimal `json:"btc"`
	Eth decimal.Decimal `json:"eth"`
	Usd decimal.Decimal `json:"usd"`
}

type Image struct {
	Thumb string `json:"thumb"`
	Small string `json:"small"`
	Large string `json:"large"`
}

type Localization struct {
	En  string `json:"en"`
	De  string `json:"de"`
	Es  string `json:"es"`
	Fr  string `json:"fr"`
	It  string `json:"it"`
	Pl  string `json:"pl"`
	Ro  string `json:"ro"`
	Hu  string `json:"hu"`
	Nl  string `json:"nl"`
	Pt  string `json:"pt"`
	Sv  string `json:"sv"`
	Vi  string `json:"vi"`
	Tr  string `json:"tr"`
	Ru  string `json:"ru"`
	Ja  string `json:"ja"`
	Zh  string `json:"zh"`
	Zhw string `json:"zhw"`
	Ko  string `json:"ko"`
	Ar  string `json:"ar"`
	Th  string `json:"th"`
	ID  string `json:"id"`
}

type CommunityData struct {
	FacebookLikes                decimal.Decimal `json:"facebook_likes"`
	TwitterFollowers             decimal.Decimal `json:"twitter_followers"`
	RedditAveragePosts           decimal.Decimal `json:"reddit_average_posts_48h"`
	RedditAverageCommentsTwoDays decimal.Decimal `json:"reddit_average_comments_48h"`
	RedditSubscribers            decimal.Decimal `json:"reddit_subscribers"`
	RedditAccountsActiveTwoDays  decimal.Decimal `json:"reddit_accounts_active_48h"`
	TelegramChannelUserCount     decimal.Decimal `json:"telegram_channel_user_count"`
}

type DeveloperData struct {
	Forks                       float32   `json:"forks"`
	Stars                       float32   `json:"stars"`
	Subscribers                 float32   `json:"subscribers"`
	TotalIssues                 float32   `json:"total_issues"`
	ClosedIssues                float32   `json:"closed_issues"`
	PullRequestsMerged          float32   `json:"pull_requests_merged"`
	PullRequestContributors     float32   `json:"pull_request_contributors"`
	CodeAdditionsDeletionsMonth CodeStats `json:"code_additions_deletions_4_weeks"`
	CommitCountMonth            float32   `json:"commit_count_4_weeks"`
	MonthCommitActivitySeries   []float32 `json:"last_4_weeks_commit_activity_series"`
}

type CodeStats struct {
	Additions int16 `json:"additions"`
	Deletions int16 `json:"deletions"`
}

type PublicInterestStats struct {
	AlexaRank decimal.Decimal `json:"alexa_rank"`
	// bing_matches `json:"bing_matches"`
}

type MarketData struct {
	CurrentPrice                          PriceRates      `json:"current_price"`
	Roi                                   Roi             `json:"roi"`
	Ath                                   PriceRates      `json:"ath"`
	AthChangePercentage                   PriceRates      `json:"ath_change_percentage"`
	AthDate                               DateRates       `json:"ath_date"`
	Atl                                   PriceRates      `json:"atl"`
	AtlChangePercentage                   PriceRates      `json:"atl_change_percentage"`
	AtlDate                               DateRates       `json:"atl_date"`
	MarketCap                             PriceRates      `json:"market_cap"`
	MarketCapRank                         int16           `json:"market_cap_rank"`
	FullyDilutedValuation                 PriceRates      `json:"fully_diluted_valuation"`
	TotalVolume                           PriceRates      `json:"total_volume"`
	High24                                PriceRates      `json:"high_24h"`
	Low24                                 PriceRates      `json:"low_24h"`
	PriceChangeDay                        decimal.Decimal `json:"price_change_24h"`
	PriceChangePercentDay                 decimal.Decimal `json:"price_change_percentage_24h"`
	PriceChangePercentWeek                decimal.Decimal `json:"price_change_percentage_7d"`
	PriceChangePercentTwoWeeks            decimal.Decimal `json:"price_change_percentage_14d"`
	PriceChangePercentMonth               decimal.Decimal `json:"price_change_percentage_30d"`
	PriceChangePercentTwoMonths           decimal.Decimal `json:"price_change_percentage_60d"`
	PriceChangePercentHalfYear            decimal.Decimal `json:"price_change_percentage_200d"`
	PriceChangePercentYear                decimal.Decimal `json:"price_change_percentage_1y"`
	MarketCapChangeDay                    decimal.Decimal `json:"market_cap_change_24h"`
	MarketCapChangePercentDay             decimal.Decimal `json:"market_cap_change_percentage_24h"`
	MarketCapChangeDayRates               PriceRates      `json:"price_change_24h_in_currency"`
	PriceChangePercentageHourCurrency     PriceRates      `json:"price_change_percentage_1h_in_currency"`
	PriceChangePercentageDayCurrency      PriceRates      `json:"price_change_percentage_24h_in_currency"`
	PriceChangePercentageWeekCurrency     PriceRates      `json:"price_change_percentage_7d_in_currency"`
	PriceChangePercentageTwoWeekCurrency  PriceRates      `json:"price_change_percentage_14d_in_currency"`
	PriceChangePercentageMonthCurrency    PriceRates      `json:"price_change_percentage_30d_in_currency"`
	PriceChangePercentageTwoMonthCurrency PriceRates      `json:"price_change_percentage_60d_in_currency"`
	PriceChangePercentageHalfYearCurrency PriceRates      `json:"price_change_percentage_200d_in_currency"`
	PriceChangePercentageYearCurrency     PriceRates      `json:"price_change_percentage_1y_in_currency"`
	MarketCapChangeDayCurrency            PriceRates      `json:"market_cap_change_24h_in_currency"`
	MarketCapChangePercentageDayCurrency  PriceRates      `json:"market_cap_change_percentage_24h_in_currency"`
	TotalSupply                           decimal.Decimal `json:"total_supply"`
	MaxSupply                             decimal.Decimal `json:"max_supply"`
	CirculatingSupply                     decimal.Decimal `json:"circulating_supply"`
	Sparkline                             Sparkline       `json:"sparkline_7d"`
	LastUpdated                           string          `json:"last_updated"`
}
type PriceRates struct {
	Aed  decimal.Decimal `json:"aed"`
	Ars  decimal.Decimal `json:"ars"`
	Aud  decimal.Decimal `json:"aud"`
	Bch  decimal.Decimal `json:"bch"`
	Bdt  decimal.Decimal `json:"bdt"`
	Bhd  decimal.Decimal `json:"bhd"`
	Bmd  decimal.Decimal `json:"bmd"`
	Bnb  decimal.Decimal `json:"bnb"`
	Brl  decimal.Decimal `json:"brl"`
	Btc  decimal.Decimal `json:"btc"`
	Cad  decimal.Decimal `json:"cad"`
	Chf  decimal.Decimal `json:"chf"`
	Clp  decimal.Decimal `json:"clp"`
	Cny  decimal.Decimal `json:"cny"`
	Czk  decimal.Decimal `json:"czk"`
	Dkk  decimal.Decimal `json:"dkk"`
	Dot  decimal.Decimal `json:"dot"`
	Eos  decimal.Decimal `json:"eos"`
	Eth  decimal.Decimal `json:"eth"`
	Eur  decimal.Decimal `json:"eur"`
	Gbp  decimal.Decimal `json:"gbp"`
	Hkd  decimal.Decimal `json:"hkd"`
	Huf  decimal.Decimal `json:"huf"`
	Idr  decimal.Decimal `json:"idr"`
	Ils  decimal.Decimal `json:"ils"`
	Inr  decimal.Decimal `json:"inr"`
	Jpy  decimal.Decimal `json:"jpy"`
	Krw  decimal.Decimal `json:"krw"`
	Kwd  decimal.Decimal `json:"kwd"`
	Lkr  decimal.Decimal `json:"lkr"`
	Ltc  decimal.Decimal `json:"ltc"`
	Mmk  decimal.Decimal `json:"mmk"`
	Mxn  decimal.Decimal `json:"mxn"`
	Myr  decimal.Decimal `json:"myr"`
	Ngn  decimal.Decimal `json:"ngn"`
	Nok  decimal.Decimal `json:"nok"`
	Nzd  decimal.Decimal `json:"nzd"`
	Php  decimal.Decimal `json:"php"`
	Pkr  decimal.Decimal `json:"pkr"`
	Pln  decimal.Decimal `json:"pln"`
	Rub  decimal.Decimal `json:"rub"`
	Sar  decimal.Decimal `json:"sar"`
	Sek  decimal.Decimal `json:"sek"`
	Sgd  decimal.Decimal `json:"sgd"`
	Thb  decimal.Decimal `json:"thb"`
	Try  decimal.Decimal `json:"try"`
	Twd  decimal.Decimal `json:"twd"`
	Uah  decimal.Decimal `json:"uah"`
	Usd  decimal.Decimal `json:"usd"`
	Vef  decimal.Decimal `json:"vef"`
	Vnd  decimal.Decimal `json:"vnd"`
	Xag  decimal.Decimal `json:"xag"`
	Xau  decimal.Decimal `json:"xau"`
	Xdr  decimal.Decimal `json:"xdr"`
	Xlm  decimal.Decimal `json:"xlm"`
	Xrp  decimal.Decimal `json:"xrp"`
	Yfi  decimal.Decimal `json:"yfi"`
	Zar  decimal.Decimal `json:"zar"`
	Bits decimal.Decimal `json:"bits"`
	Link decimal.Decimal `json:"link"`
	Sats decimal.Decimal `json:"sats"`
}

type DateRates struct {
	Aed  string `json:"aed"`
	Ars  string `json:"ars"`
	Aud  string `json:"aud"`
	Bch  string `json:"bch"`
	Bdt  string `json:"bdt"`
	Bhd  string `json:"bhd"`
	Bmd  string `json:"bmd"`
	Bnb  string `json:"bnb"`
	Brl  string `json:"brl"`
	Btc  string `json:"btc"`
	Cad  string `json:"cad"`
	Chf  string `json:"chf"`
	Clp  string `json:"clp"`
	Cny  string `json:"cny"`
	Czk  string `json:"czk"`
	Dkk  string `json:"dkk"`
	Dot  string `json:"dot"`
	Eos  string `json:"eos"`
	Eth  string `json:"eth"`
	Eur  string `json:"eur"`
	Gbp  string `json:"gbp"`
	Hkd  string `json:"hkd"`
	Huf  string `json:"huf"`
	Idr  string `json:"idr"`
	Ils  string `json:"ils"`
	Inr  string `json:"inr"`
	Jpy  string `json:"jpy"`
	Krw  string `json:"krw"`
	Kwd  string `json:"kwd"`
	Lkr  string `json:"lkr"`
	Ltc  string `json:"ltc"`
	Mmk  string `json:"mmk"`
	Mxn  string `json:"mxn"`
	Myr  string `json:"myr"`
	Ngn  string `json:"ngn"`
	Nok  string `json:"nok"`
	Nzd  string `json:"nzd"`
	Php  string `json:"php"`
	Pkr  string `json:"pkr"`
	Pln  string `json:"pln"`
	Rub  string `json:"rub"`
	Sar  string `json:"sar"`
	Sek  string `json:"sek"`
	Sgd  string `json:"sgd"`
	Thb  string `json:"thb"`
	Try  string `json:"try"`
	Twd  string `json:"twd"`
	Uah  string `json:"uah"`
	Usd  string `json:"usd"`
	Vef  string `json:"vef"`
	Vnd  string `json:"vnd"`
	Xag  string `json:"xag"`
	Xau  string `json:"xau"`
	Xdr  string `json:"xdr"`
	Xlm  string `json:"xlm"`
	Xrp  string `json:"xrp"`
	Yfi  string `json:"yfi"`
	Zar  string `json:"zar"`
	Bits string `json:"bits"`
	Link string `json:"link"`
	Sats string `json:"sats"`
}

type CapRates struct {
	Aed  int64 `json:"aed"`
	Ars  int64 `json:"ars"`
	Aud  int64 `json:"aud"`
	Bch  int64 `json:"bch"`
	Bdt  int64 `json:"bdt"`
	Bhd  int64 `json:"bhd"`
	Bmd  int64 `json:"bmd"`
	Bnb  int64 `json:"bnb"`
	Brl  int64 `json:"brl"`
	Btc  int64 `json:"btc"`
	Cad  int64 `json:"cad"`
	Chf  int64 `json:"chf"`
	Clp  int64 `json:"clp"`
	Cny  int64 `json:"cny"`
	Czk  int64 `json:"czk"`
	Dkk  int64 `json:"dkk"`
	Dot  int64 `json:"dot"`
	Eos  int64 `json:"eos"`
	Eth  int64 `json:"eth"`
	Eur  int64 `json:"eur"`
	Gbp  int64 `json:"gbp"`
	Hkd  int64 `json:"hkd"`
	Huf  int64 `json:"huf"`
	Idr  int64 `json:"idr"`
	Ils  int64 `json:"ils"`
	Inr  int64 `json:"inr"`
	Jpy  int64 `json:"jpy"`
	Krw  int64 `json:"krw"`
	Kwd  int64 `json:"kwd"`
	Lkr  int64 `json:"lkr"`
	Ltc  int64 `json:"ltc"`
	Mmk  int64 `json:"mmk"`
	Mxn  int64 `json:"mxn"`
	Myr  int64 `json:"myr"`
	Ngn  int64 `json:"ngn"`
	Nok  int64 `json:"nok"`
	Nzd  int64 `json:"nzd"`
	Php  int64 `json:"php"`
	Pkr  int64 `json:"pkr"`
	Pln  int64 `json:"pln"`
	Rub  int64 `json:"rub"`
	Sar  int64 `json:"sar"`
	Sek  int64 `json:"sek"`
	Sgd  int64 `json:"sgd"`
	Thb  int64 `json:"thb"`
	Try  int64 `json:"try"`
	Twd  int64 `json:"twd"`
	Uah  int64 `json:"uah"`
	Usd  int64 `json:"usd"`
	Vef  int64 `json:"vef"`
	Vnd  int64 `json:"vnd"`
	Xag  int64 `json:"xag"`
	Xau  int64 `json:"xau"`
	Xdr  int64 `json:"xdr"`
	Xlm  int64 `json:"xlm"`
	Xrp  int64 `json:"xrp"`
	Yfi  int64 `json:"yfi"`
	Zar  int64 `json:"zar"`
	Bits int64 `json:"bits"`
	Link int64 `json:"link"`
	Sats int64 `json:"sats"`
}
