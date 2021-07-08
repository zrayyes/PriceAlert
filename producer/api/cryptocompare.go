package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Return a JSON map of crypto coin prices in the given currencies
func GetCoinPrices(coins []string, currencies []string) map[string]map[string]float64 {
	var results map[string]map[string]float64

	resp, err := http.Get(fmt.Sprintf("https://min-api.cryptocompare.com/data/pricemulti?fsyms=%s&tsyms=%s", strings.Join(coins, ","), strings.Join(currencies, ",")))
	if err != nil {
		fmt.Println("GetCoinPrices: Cant connect to API: ", err.Error())
		return results
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("GetCoinPrices: Cant read body: ", err.Error())
		return results
	}
	json.Unmarshal([]byte(body), &results)
	return results
}
