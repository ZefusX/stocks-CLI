package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/fatih/color"
)

// vGPHQNGIth0E33iJk1R7l6epTGSBxQLa

type Stock struct {
	MetaData struct {
		Symbol    string `json:"2. Symbol"`
		Refreshed string `json:"3. Last Refreshed"`
	} `json:"Meta Data"`
	TimeSeries map[string]struct {
		Open   string `json:"1. open"`
		High   string `json:"2. high"`
		Low    string `json:"3. low"`
		Close  string `json:"4. close"`
		Volume string `json:"5. volume"`
	} `json:"Time Series (5min)"`
}

func main() {
	q := ""
	if len(os.Args) >= 2 {
		q = os.Args[1]
	} else {
		panic("Syntax : stock *Symbol*, only one argument")
	}

	response, err := http.Get("https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=" + q + "&interval=5min&outputsize=full&apikey=SZPMBD2PGDOKGUJ3")
	if err != nil {
		panic(err)
	}

	// Ferme la réponse
	defer response.Body.Close()
	if response.StatusCode != 200 {
		panic("Stock API not available")
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var stock Stock
	// Pour récupérer le json et formater avec stock
	err = json.Unmarshal(body, &stock)
	if err != nil {
		panic(err)
	}

	metadata := stock.MetaData
	fmt.Printf("%s, %s\n", metadata.Symbol, metadata.Refreshed)

	// Accéder au dernier élément de Time Series
	lastTime := ""
	for timestamp := range stock.TimeSeries {
		lastTime = timestamp
	}
	lastData := stock.TimeSeries[lastTime]

	c := color.New(color.FgGreen)
	c.Printf("-----%s-----\nLast Refreshed (%s):\nOpen: %s\nHigh: %s\nLow: %s\nClose: %s\nVolume: %s\n ", metadata.Symbol, metadata.Refreshed, lastData.Open, lastData.High, lastData.Low, lastData.Close, lastData.Volume)
}
