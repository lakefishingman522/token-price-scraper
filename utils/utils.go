package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/CascadiaFoundation/CascadiaTokenScrapper/models"
)

type CryptoCompareResponse struct {
	Response string `json:"Response"`
	Message  string `json:"Message"`
	Data     struct {
		Aggregated bool `json:"Aggregated"`
		TimeFrom   int  `json:"TimeFrom"`
		TimeTo     int  `json:"TimeTo"`
		Data       []struct {
			Time             int     `json:"time"`
			High             float64 `json:"high"`
			Low              float64 `json:"low"`
			Open             float64 `json:"open"`
			VolumeFrom       float64 `json:"volumefrom"`
			VolumeTo         float64 `json:"volumeto"`
			Close            float64 `json:"close"`
			ConversionType   string  `json:"conversionType"`
			ConversionSymbol string  `json:"conversionSymbol"`
		} `json:"Data"`
	} `json:"Data"`
}

type HTTPError struct {
	Message string `json:"message"`
}

type TokenStatistics struct {
	PreviousAveragePrice float64 `json:"previous_avg_price"`
	CurrentAveragePrice  float64 `json:"current_avg_price"`
	PercentageChange     float64 `json:"percentage_change"`
}

func GetPriceStatistics() models.TokenStatisticsModel {
	client := resty.New()

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"fsym":    "ETH", // from symbol
			"tsym":    "USD", // to symbol
			"limit":   "720", // last 720 days
			"api_key": "ae3b6f305c759ce21aa7d9e080566db8a4bf52a4a6d3a1eeae995ee4c6b0d3e6",
		}).
		Get("https://min-api.cryptocompare.com/data/v2/histoday")

	if err != nil {
		log.Fatalf("Error occurred while making the request: %v", err)
	}

	var result CryptoCompareResponse

	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		log.Fatalf("Error occurred while unmarshalling the response: %v", err)
	}

	fmt.Printf("Average ETH price change during 360 days: $%.2f\n", CalPercentageChange(360, result))
	fmt.Printf("Average ETH price change during 180 days: $%.2f\n", CalPercentageChange(180, result))
	fmt.Printf("Average ETH price change during  90 days: $%.2f\n", CalPercentageChange(90, result))
	fmt.Printf("Average ETH price change during  30 days: $%.2f\n", CalPercentageChange(30, result))
	fmt.Printf("Average ETH price change during  14 days: $%.2f\n", CalPercentageChange(14, result))
	fmt.Printf("Average ETH price change during   7 days: $%.2f\n", CalPercentageChange(7, result))

	// =============================================== Get Avg Price of previous 48 hours to calculate pricePercentageChange =================
	resp_, err_ := client.R().
		SetQueryParams(map[string]string{
			"fsym":    "ETH", // from symbol
			"tsym":    "USD", // to symbol
			"limit":   "48",  // last 48 hours
			"api_key": "ae3b6f305c759ce21aa7d9e080566db8a4bf52a4a6d3a1eeae995ee4c6b0d3e6",
		}).
		Get("https://min-api.cryptocompare.com/data/v2/histohour")

	if err_ != nil {
		log.Fatalf("Error occurred while making the 48 hours request: %v", err)
	}

	var result_ CryptoCompareResponse

	err_ = json.Unmarshal(resp_.Body(), &result_)
	if err_ != nil {
		log.Fatalf("Error occurred while unmarshalling the 48 hours response: %v", err)
	}

	fmt.Printf("Average ETH price change during   1 days: $%.2f\n", CalPercentageChange(24, result))

	tStats_360 := CalPercentageChange(360, result)
	tStats_180 := CalPercentageChange(180, result)
	tStats_90 := CalPercentageChange(90, result)
	tStats_30 := CalPercentageChange(30, result)
	tStats_14 := CalPercentageChange(14, result)
	tStats_7 := CalPercentageChange(7, result)
	tStats_1 := CalPercentageChange(24, result_)

	return models.TokenStatisticsModel{
		TimeStamp: uint64(time.Now().Unix()),

		P360: tStats_360.PercentageChange,
		P180: tStats_180.PercentageChange,
		P90:  tStats_90.PercentageChange,
		P30:  tStats_30.PercentageChange,
		P14:  tStats_14.PercentageChange,
		P7:   tStats_7.PercentageChange,
		P1:   tStats_1.PercentageChange,

		PreviousAveragePrice360: tStats_360.PreviousAveragePrice,
		PreviousAveragePrice180: tStats_180.PreviousAveragePrice,
		PreviousAveragePrice90:  tStats_90.PreviousAveragePrice,
		PreviousAveragePrice30:  tStats_30.PreviousAveragePrice,
		PreviousAveragePrice14:  tStats_14.PreviousAveragePrice,
		PreviousAveragePrice7:   tStats_7.PreviousAveragePrice,
		PreviousAveragePrice1:   tStats_1.PreviousAveragePrice,

		CurrentAveragePrice360: tStats_360.CurrentAveragePrice,
		CurrentAveragePrice180: tStats_180.CurrentAveragePrice,
		CurrentAveragePrice90:  tStats_90.CurrentAveragePrice,
		CurrentAveragePrice30:  tStats_30.CurrentAveragePrice,
		CurrentAveragePrice14:  tStats_14.CurrentAveragePrice,
		CurrentAveragePrice7:   tStats_7.CurrentAveragePrice,
		CurrentAveragePrice1:   tStats_1.CurrentAveragePrice,
	}
}

func CalPercentageChange(period int, respData CryptoCompareResponse) TokenStatistics {
	length := len(respData.Data.Data)

	// Calculate average of previous Period days
	var sumPreviousPeriod float64
	for i := length - 2*period; i < length-period; i++ {
		sumPreviousPeriod += respData.Data.Data[i].Close
	}

	averagePreviousPeriod := math.Round(sumPreviousPeriod/float64(period)*100) / 100

	// Calculate average of last Period days
	var sumLastPeriod float64
	for i := length - period; i < length; i++ {
		sumLastPeriod += respData.Data.Data[i].Close
	}

	averageLastPeriod := math.Round(sumLastPeriod/float64(period)*100) / 100

	return TokenStatistics{
		PreviousAveragePrice: averagePreviousPeriod,
		CurrentAveragePrice:  averageLastPeriod,
		PercentageChange:     averageLastPeriod / averagePreviousPeriod,
	}

}
