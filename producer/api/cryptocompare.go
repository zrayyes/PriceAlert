package api

import (
	"encoding/json"
)

// https://min-api.cryptocompare.com/data/pricemulti?fsyms=BTC&tsyms=USD&api_key=

func GetCoinPrices(coins []string, currencies []string) map[string]map[string]float64 {
	var results map[string]map[string]float64
	json.Unmarshal([]byte(`{"BTC":{"USD":34554.5}}`), &results)
	return results
}
