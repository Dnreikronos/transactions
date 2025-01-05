package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func FetchExchangeRates(currencyDescription string, transactionDate time.Time) (float64, error) {

	formattedDate := transactionDate.Format("2006-01-02")

	filter := fmt.Sprintf("record_date:lte:%s,country_currency_desc:eq:%s", formattedDate, currencyDescription)
	fields := "exchange_rate"

	requestURL := fmt.Sprintf(
		"https://api.fiscaldata.treasury.gov/services/api/fiscal_service/v1/accounting/od/rates_of_exchange?filter=%s&fields=%s",
		filter, fields,
	)

	resp, err := http.Get(requestURL)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch exchange rates: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		responseBody, _ := io.ReadAll(resp.Body)
		return 0, fmt.Errorf("API error: %d - %s", resp.StatusCode, string(responseBody))
	}

	var response struct {
		Data []struct {
			ExchangeRate        string `json:"exchange_rate"`
			RecordDate          string `json:"record_date"`
			CountryCurrencyDesc string `json:"country_currency_desc"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return 0, fmt.Errorf("failed to decode API response: %v", err)
	}

	fmt.Printf("Decoded API Response: %+v\n", response)

	if len(response.Data) == 0 {
		return 0, fmt.Errorf("no exchange rate data available for currency: %s on date: %s", currencyDescription, formattedDate)
	}

	for _, record := range response.Data {

		fmt.Printf("Comparing currency: '%s' with record: '%s'\n", currencyDescription, record.CountryCurrencyDesc)

		if strings.TrimSpace(record.CountryCurrencyDesc) == strings.TrimSpace(currencyDescription) && record.RecordDate == formattedDate {

			exchangeRate, err := strconv.ParseFloat(record.ExchangeRate, 64)
			if err != nil {
				return 0, fmt.Errorf("failed to parse exchange rate: %v", err)
			}

			return exchangeRate, nil
		}
	}

	return 0, fmt.Errorf("no exchange rate found for currency: %s on date: %s", currencyDescription, formattedDate)
}
