package main

import (
	"fmt"

	gecko "github.com/verzth/go-coingecko"
)

func main() {
	cgClient := gecko.NewClient(nil)
	defer cgClient.Close()

	data, err := cgClient.CoinsId("bitcoin", true, true, true, false, false, false)
	if err != nil {
		fmt.Print("Somethig went wrong...")
		return
	}
	bitcoinPrice, _ := data.MarketData.CurrentPrice.Usd.Float64()
	fmt.Printf("Bitcoin price is: %f$", bitcoinPrice)
}
