package main

import (
	"encoding/json"
	"strconv"
)

// List of price estimates
type PriceEstimates struct {
	StartLatitude  float64
	StartLongitude float64
	EndLatitude    float64
	EndLongitude   float64
	Prices         []PriceEstimate `json:"prices"`
}

// Uber price estimate
type PriceEstimate struct {
	ProductId       string  `json:"product_id"`
	CurrencyCode    string  `json:"currency_code"`
	DisplayName     string  `json:"display_name"`
	Estimate        string  `json:"estimate"`
	LowEstimate     int     `json:"low_estimate"`
	HighEstimate    int     `json:"high_estimate"`
	SurgeMultiplier float64 `json:"surge_multiplier"`
	Duration        int     `json:"duration"`
	Distance        float64 `json:"distance"`
}

// Internal method that implements the Getter interface
func (pe *PriceEstimates) get(c *Client) error {
	priceEstimateParams := map[string]string{
		"start_latitude":  strconv.FormatFloat(pe.StartLatitude, 'f', 2, 32),
		"start_longitude": strconv.FormatFloat(pe.StartLongitude, 'f', 2, 32),
		"end_latitude":    strconv.FormatFloat(pe.EndLatitude, 'f', 2, 32),
		"end_longitude":   strconv.FormatFloat(pe.EndLongitude, 'f', 2, 32),
	}

	data := c.getRequest("estimates/price", priceEstimateParams)
	if e := json.Unmarshal(data, &pe); e != nil {
		return e
	}
	return nil
}