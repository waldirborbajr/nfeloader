package services

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
)

type QuotationService struct {
	HTTPClient *http.Client
}

type currencyData struct {
	Chart struct {
		Result []struct {
			Meta struct {
				Currency             string  `json:"currency"`
				Symbol               string  `json:"symbol"`
				ExchangeName         string  `json:"exchangeName"`
				InstrumentType       string  `json:"instrumentType"`
				FirstTradeDate       int     `json:"firstTradeDate"`
				RegularMarketTime    int     `json:"regularMarketTime"`
				Gmtoffset            int     `json:"gmtoffset"`
				Timezone             string  `json:"timezone"`
				ExchangeTimezoneName string  `json:"exchangeTimezoneName"`
				RegularMarketPrice   float64 `json:"regularMarketPrice"`
				ChartPreviousClose   float64 `json:"chartPreviousClose"`
				PreviousClose        float64 `json:"previousClose"`
				Scale                int     `json:"scale"`
				PriceHint            int     `json:"priceHint"`
				CurrentTradingPeriod struct {
					Pre struct {
						Timezone  string `json:"timezone"`
						Start     int    `json:"start"`
						End       int    `json:"end"`
						Gmtoffset int    `json:"gmtoffset"`
					} `json:"pre"`
					Regular struct {
						Timezone  string `json:"timezone"`
						Start     int    `json:"start"`
						End       int    `json:"end"`
						Gmtoffset int    `json:"gmtoffset"`
					} `json:"regular"`
					Post struct {
						Timezone  string `json:"timezone"`
						Start     int    `json:"start"`
						End       int    `json:"end"`
						Gmtoffset int    `json:"gmtoffset"`
					} `json:"post"`
				} `json:"currentTradingPeriod"`
				TradingPeriods [][]struct {
					Timezone  string `json:"timezone"`
					Start     int    `json:"start"`
					End       int    `json:"end"`
					Gmtoffset int    `json:"gmtoffset"`
				} `json:"tradingPeriods"`
				DataGranularity string   `json:"dataGranularity"`
				Range           string   `json:"range"`
				ValidRanges     []string `json:"validRanges"`
			} `json:"meta"`
			Timestamp  []int `json:"timestamp"`
			Indicators struct {
				Quote []struct {
					Volume []int     `json:"volume"`
					Low    []float64 `json:"low"`
					Open   []float64 `json:"open"`
					Close  []float64 `json:"close"`
					High   []float64 `json:"high"`
				} `json:"quote"`
			} `json:"indicators"`
		} `json:"result"`
		Error interface{} `json:"error"`
	} `json:"chart"`
}

func NewQuotationService(httpClient *http.Client) *QuotationService {
	return &QuotationService{
		HTTPClient: httpClient,
	}
}

func (quotationService QuotationService) GetQuote(symbol string) (*float32, error) {

	url := "https://query1.finance.yahoo.com/v8/finance/chart/USDBRL=X?region=US&lang=en-US&includePrePost=false&interval=2m&useYfid=true&range=1d&corsDomain=finance.yahoo.com&.tsrc=finance"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Host", "query1.finance.yahoo.com")
	req.Header.Add("Accept-Language", "en-GB,en;q=0.9")
	req.Header.Add("Referer", "https://finance.yahoo.com/quote/usdbrl=x/")
	req.Header.Add("Connection", "keep-alive")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	// Check that the server actually sent compressed data
	var reader io.ReadCloser
	reader, err = gzip.NewReader(res.Body)
	if err != nil {
		return nil, err
	}

	var jsonData currencyData
	err = json.NewDecoder(reader).Decode(&jsonData)
	if err != nil {
		return nil, err
	}

	defer reader.Close()

	price := float32(jsonData.Chart.Result[0].Meta.RegularMarketPrice)

	return &price, nil

}
