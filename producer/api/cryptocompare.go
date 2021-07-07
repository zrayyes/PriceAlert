package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetCoinPrices(coins []string, currencies []string) map[string]map[string]float64 {
	var results map[string]map[string]float64
	// body := `{"BTC":{"USD":34554.5},"ETC":{"EUR":45.55}}`

	resp, err := http.Get(fmt.Sprintf("https://min-api.cryptocompare.com/data/pricemulti?fsyms=%s&tsyms=%s", strings.Join(coins, ","), strings.Join(currencies, ",")))
	if err != nil {
		log.Fatalln("Cant connect to API.")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	json.Unmarshal([]byte(body), &results)
	return results
}
