package nse

import (
	"encoding/json"
	"errors"
	"log"
	"net/url"
	"os"
	"slices"
	"strings"

	"github.com/go-resty/resty/v2"
)

const apiURL = "https://www.nseindia.com"

var (
	baseHeaders = map[string]string{
		"Accept-Language": "en-US,en;q=0.9",
		"Accept-Encoding": "gzip, deflate, br",
		"Connection":      "keep-alive",
		"User-Agent":      "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/118.0",
	}

	client = initRestyClient(apiURL, baseHeaders)
)

// initializeRestyClient initializes and returns a resty.Client with the provided base URL and headers
func initRestyClient(baseURL string, headers map[string]string) *resty.Client {
	return resty.New().SetBaseURL(baseURL).SetHeaders(headers)
}

// getCookie obtains the required cookies from the NSE website's home page
func getCookie() string {
	response, err := client.R().EnableTrace().Get("/")
	if err != nil {
		log.Fatal("Failed to get cookie:", err)
	}

	cookies := response.Cookies()
	var cook []string

	requiredCookies := []string{"nsit", "nseappid", "ak_bmsc", "AKA_A2", "bm_mi", "bm_sv"}
	for _, cookie := range cookies {
		keyVal := strings.Split(cookie.String(), ";")[0]
		key := strings.Split(keyVal, "=")[0]
		if slices.Contains(requiredCookies, key) {
			cook = append(cook, keyVal)
		}
	}

	return strings.Join(cook, "; ")
}

// marketDataPreOpen fetches market data for pre-open
func MarketDataPreOpen() (*StockData, error) {
	cookie := getCookie()
	response, err := client.R().EnableTrace().SetHeader("Cookie", cookie).Get("/api/market-data-pre-open?key=ALL")
	if err != nil {
		log.Fatal("Failed to fetch market data:", err)
	}

	if response.StatusCode() == 200 {
		var stockData StockData
		err := json.Unmarshal(response.Body(), &stockData)
		if err != nil {
			log.Println("Error decoding market data:", err)
			return nil, err
		}
		return &stockData, nil
	}
	return nil, errors.New("failed to fetch market data")
}

// getSymbols retrieves symbols from market data
func GetSymbols() []string {
	res, err := MarketDataPreOpen()
	if err != nil {
		log.Println("Error getting symbols:", err)
	}
	var symbols []string
	for _, val := range res.Data {
		symbols = append(symbols, val.Metadata.Symbol)
	}
	return symbols
}

// quoteEquity fetches equity details for a given symbol
func QuoteEquity(symbol string) (*EquityDetails, error) {
	cookie := getCookie()
	response, err := client.R().EnableTrace().SetHeader("Cookie", cookie).
		Get("/api/quote-equity?symbol=" + url.QueryEscape(strings.ToUpper(symbol)))
	if err != nil {
		log.Fatal("Failed to fetch equity details:", err)
	}
	if response.StatusCode() == 200 {
		var stockData EquityDetails
		// os.WriteFile("TATATECH.json", response.Body(), 0644)
		err := json.Unmarshal(response.Body(), &stockData)
		if err != nil {
			log.Println("Error decoding equity details:", err)
			return nil, err
		}
		return &stockData, nil
	}
	return nil, errors.New("failed to fetch equity details")
}

func QuoteEquityTradeInfo(symbol string) (*EquityTradeInfo, error) {
	cookie := getCookie()
	response, err := client.R().EnableTrace().SetHeader("Cookie", cookie).
		Get("/api/quote-equity?symbol=" + url.QueryEscape(strings.ToUpper(symbol)) + "&section=trade_info")
	if err != nil {
		log.Fatal("Failed to fetch equity details:", err)
	}
	if response.StatusCode() == 200 {
		var stockData EquityTradeInfo
		os.WriteFile("MITCON.json", response.Body(), 0644)
		err := json.Unmarshal(response.Body(), &stockData)
		if err != nil {
			log.Println("Error decoding equity details:", err)
			return nil, err
		}
		return &stockData, nil
	}
	return nil, errors.New("failed to fetch equity details")
}

func ChartDataByIndexPreopen(symbol string) (*IntradayData, error) {
	details, _ := QuoteEquity(symbol)
	identifier := details.Info.Identifier
	cookie := getCookie()
	url := "/api/chart-databyindex?index=" + url.QueryEscape(identifier) + "&preopen=true"
	response, err := client.R().EnableTrace().SetHeader("Cookie", cookie).
		Get(url)
	if err != nil {
		log.Fatal("Failed to fetch equity details:", err)
	}
	if response.StatusCode() == 200 {
		var stockData IntradayData
		os.WriteFile("MITCON.json", response.Body(), 0644)
		err := json.Unmarshal(response.Body(), &stockData)
		if err != nil {
			log.Println("Error decoding equity details:", err)
			return nil, err
		}
		return &stockData, nil
	}
	return nil, errors.New("failed to fetch equity details")
}

func ChartDataByIndex(symbol string) (*IntradayData, error) {
	details, _ := QuoteEquity(symbol)
	identifier := details.Info.Identifier
	cookie := getCookie()
	url := "/api/chart-databyindex?index=" + url.QueryEscape(identifier)
	response, err := client.R().EnableTrace().SetHeader("Cookie", cookie).
		Get(url)
	if err != nil {
		log.Fatal("Failed to fetch equity details:", err)
	}
	if response.StatusCode() == 200 {
		var stockData IntradayData
		os.WriteFile("MITCON.json", response.Body(), 0644)
		err := json.Unmarshal(response.Body(), &stockData)
		if err != nil {
			log.Println("Error decoding equity details:", err)
			return nil, err
		}
		return &stockData, nil
	}
	return nil, errors.New("failed to fetch equity details")
}

// lol is a sample function for testing
func Lol() {
	val, err := ChartDataByIndex("MITCON")
	if err != nil {
		log.Println("Error in lol:", err)
		return
	}
	log.Printf("%v", val)
}
