package nse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitRestyClient(t *testing.T) {
	const apiURL = "https://www.nseindia.com"
	baseHeaders = map[string]string{
		"Accept-Language": "en-US,en;q=0.9",
		"Accept-Encoding": "gzip, deflate, br",
		"Connection":      "keep-alive",
		"User-Agent":      "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/118.0",
	}

	client = initRestyClient(apiURL, baseHeaders)
	assert.Equal(t, apiURL, client.HostURL)
	assert.Equal(t, "en-US,en;q=0.9", client.Header.Get("Accept-Language"))
}

func TestGetCookie(t *testing.T) {
	cookie := getCookie()
	assert.NotEmpty(t, cookie)
}

func TestMarketDataPreOpen(t *testing.T) {

	stockData, err := MarketDataPreOpen()
	assert.NoError(t, err)
	assert.NotNil(t, stockData)
	// data, _ := json.Marshal(stockData)
	// os.WriteFile("marketDataPreOpen.json", data, 0644)
	assert.NotEmpty(t, stockData.Data)
}

func TestGetSymbols(t *testing.T) {
	symbols := GetSymbols()
	assert.NotEmpty(t, symbols)
	assert.Contains(t, symbols, "ZEEMEDIA")
	assert.Contains(t, symbols, "TATATECH")
}

func TestQuoteEquity(t *testing.T) {
	var stockData *EquityDetails
	stockData, err := QuoteEquity("MITCON")
	assert.NoError(t, err)
	assert.NotNil(t, stockData)
	// log.Printf("Response: %+v\n", stockData)
	assert.Equal(t, "MITCON", stockData.Info.Symbol)
}
