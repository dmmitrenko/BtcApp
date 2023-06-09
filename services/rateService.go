package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"example/BtcApp/models"
	"example/BtcApp/utils"
)

func FetchExchangeRate() (models.Rate, error) {
	url := utils.CoingeckoAPIURL

	resp, err := http.Get(url)
	if err != nil {
		return models.Rate{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.Rate{}, err
	}

	var exchangeRates map[string]models.Rate
	err = json.Unmarshal(body, &exchangeRates)
	if err != nil {
		return models.Rate{}, err
	}

	hryvniaRate, ok := exchangeRates["uah"]
	if !ok {
		return models.Rate{}, fmt.Errorf("Exchange rate for UAH not found")
	}

	return hryvniaRate, nil
}
