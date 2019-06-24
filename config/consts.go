package config

const (
	Env_DEVELOPMENT string = "development"
	Env_AMAZON_TEST string = "amazon_test"
	Env_AMAZON_AWS  string = "amazon_aws"
)

const (
	CoinId_USDT int = 1
	CoinId_BTC  int = 2
	CoinId_ETH  int = 3
)

var CoinName = map[int]string{
	CoinId_USDT: "USDT",
	CoinId_BTC:  "BTC",
	CoinId_ETH:  "ETH",
}
