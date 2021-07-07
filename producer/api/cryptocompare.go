package api

import (
	"encoding/json"
	"fmt"
	"strings"
)

func GetCoinPrices(coins []string, currencies []string) map[string]map[string]float64 {
	var results map[string]map[string]float64
	fmt.Printf("https://min-api.cryptocompare.com/data/pricemulti?fsyms=%s&tsyms=%s&api_key=\n", strings.Join(coins, ","), strings.Join(currencies, ","))
	json.Unmarshal([]byte(`{"BTC":{"USD":34554.5}}`), &results)
	return results
}
