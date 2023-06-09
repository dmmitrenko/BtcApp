package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"BtcApp/models"
	"BtcApp/utils"
)

type ExchangeRateResponse struct {
	Rates map[string]models.Rate `json:"rates"`
}

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

	var exchangeRates ExchangeRateResponse
	err = json.Unmarshal(body, &exchangeRates)
	if err != nil {
		return models.Rate{}, err
	}

	hryvniaRate := exchangeRates.Rates["uah"]

	return hryvniaRate, nil
}
