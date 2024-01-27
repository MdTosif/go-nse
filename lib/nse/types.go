package nse

import "time"

// Metadata represents metadata information in the JSON
type Metadata struct {
	Symbol         string  `json:"symbol"`
	Identifier     string  `json:"identifier"`
	Purpose        string  `json:"purpose"`
	LastPrice      float64 `json:"lastPrice"`
	Change         float64 `json:"change"`
	PChange        float64 `json:"pChange"`
	PreviousClose  float64 `json:"previousClose"`
	FinalQuantity  int     `json:"finalQuantity"`
	TotalTurnover  float64 `json:"totalTurnover"`
	MarketCap      string  `json:"marketCap"`
	YearHigh       float64 `json:"yearHigh"`
	YearLow        float64 `json:"yearLow"`
	Iep            float64 `json:"iep"`
	ChartTodayPath string  `json:"chartTodayPath"`
}

// Detail represents detailed information in the JSON
type Detail struct {
	PreOpenMarket struct {
		Preopen []struct {
			Price   float64 `json:"price"`
			BuyQty  int     `json:"buyQty"`
			SellQty int     `json:"sellQty"`
		} `json:"preopen"`
		Ato struct {
			TotalBuyQuantity  int `json:"totalBuyQuantity"`
			TotalSellQuantity int `json:"totalSellQuantity"`
		} `json:"ato"`
		Iep               float64 `json:"IEP"`
		TotalTradedVolume int     `json:"totalTradedVolume"`
		FinalPrice        float64 `json:"finalPrice"`
		FinalQuantity     int     `json:"finalQuantity"`
		LastUpdateTime    string  `json:"lastUpdateTime"`
		TotalSellQuantity int     `json:"totalSellQuantity"`
		TotalBuyQuantity  int     `json:"totalBuyQuantity"`
		AtoBuyQty         int     `json:"atoBuyQty"`
		AtoSellQty        int     `json:"atoSellQty"`
		Change            float64 `json:"Change"`
		PerChange         float64 `json:"perChange"`
		PrevClose         float64 `json:"prevClose"`
	} `json:"preOpenMarket"`
}

// StockData represents the entire JSON structure
type StockData struct {
	Data []struct {
		Metadata Metadata `json:"metadata"`
		Detail   Detail   `json:"detail"`
	} `json:"data"`
}

type IntradayData struct {
	Identifier string     `json:"identifier"`
	Name       string     `json:"name"`
	GraphData  [2]float64 `json:"graphData"`
	ClosePrice float64    `json:"closePrice"`
}

type DateRange struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

type EquityInfo struct {
	Symbol              string   `json:"symbol"`
	CompanyName         string   `json:"companyName"`
	Industry            string   `json:"industry"`
	ActiveSeries        []string `json:"activeSeries"`
	DebtSeries          []string `json:"debtSeries"`
	TempSuspendedSeries []string `json:"tempSuspendedSeries"`
	IsFNOSec            bool     `json:"isFNOSec"`
	IsCASec             bool     `json:"isCASec"`
	IsSLBSec            bool     `json:"isSLBSec"`
	IsDebtSec           bool     `json:"isDebtSec"`
	IsSuspended         bool     `json:"isSuspended"`
	IsETFSec            bool     `json:"isETFSec"`
	IsDelisted          bool     `json:"isDelisted"`
	Isin                string   `json:"isin"`
	IsTop10             bool     `json:"isTop10"`
	Identifier          string   `json:"identifier"`
}

type EquityMetadata struct {
	Series         string  `json:"series"`
	Symbol         string  `json:"symbol"`
	Isin           string  `json:"isin"`
	Status         string  `json:"status"`
	ListingDate    string  `json:"listingDate"`
	Industry       string  `json:"industry"`
	LastUpdateTime string  `json:"lastUpdateTime"`
	PdSectorPe     float64 `json:"pdSectorPe"`
	PdSymbolPe     float64 `json:"pdSymbolPe"`
	PdSectorInd    string  `json:"pdSectorInd"`
}

type EquitySecurityInfo struct {
	BoardStatus    string       `json:"boardStatus"`
	TradingStatus  string       `json:"tradingStatus"`
	TradingSegment string       `json:"tradingSegment"`
	SessionNo      string       `json:"sessionNo"`
	Slb            string       `json:"slb"`
	ClassOfShare   string       `json:"classOfShare"`
	Derivatives    string       `json:"derivatives"`
	Surveillance   Surveillance `json:"surveillance"`
	FaceValue      float64      `json:"faceValue"`
	IssuedCap      float64      `json:"issuedCap"`
	IssuedSize     float64      `json:"issuedSize"`
}

type Surveillance struct {
	Surv string `json:"surv"`
	Desc string `json:"desc"`
}

type EquityPriceInfo struct {
	LastPrice       float64 `json:"lastPrice"`
	Change          float64 `json:"change"`
	PChange         float64 `json:"pChange"`
	PreviousClose   float64 `json:"previousClose"`
	Open            float64 `json:"open"`
	Close           float64 `json:"close"`
	Vwap            float64 `json:"vwap"`
	LowerCP         string  `json:"lowerCP"`
	UpperCP         string  `json:"upperCP"`
	PPriceBand      string  `json:"pPriceBand"`
	BasePrice       float64 `json:"basePrice"`
	IntraDayHighLow struct {
		Min   float64 `json:"min"`
		Max   float64 `json:"max"`
		Value float64 `json:"value"`
	} `json:"intraDayHighLow"`
	WeekHighLow struct {
		Min     float64 `json:"min"`
		MinDate string  `json:"minDate"`
		Max     float64 `json:"max"`
		MaxDate string  `json:"maxDate"`
		Value   float64 `json:"value"`
	} `json:"weekHighLow"`
	INavValue string `json:"iNavValue"`
	CheckINAV bool   `json:"checkINAV"`
}

type PreOpenDetails struct {
	Price   float64 `json:"price"`
	BuyQty  int     `json:"buyQty"`
	SellQty int     `json:"sellQty"`
	Iep     bool    `json:"iep"`
}

type EquityPreOpenMarket struct {
	Preopen []PreOpenDetails `json:"preopen"`
	Ato     struct {
		Buy  int `json:"buy"`
		Sell int `json:"sell"`
	} `json:"ato"`
	IEP               float64 `json:"IEP"`
	TotalTradedVolume int     `json:"totalTradedVolume"`
	FinalPrice        float64 `json:"finalPrice"`
	FinalQuantity     int     `json:"finalQuantity"`
	LastUpdateTime    string  `json:"lastUpdateTime"`
	TotalBuyQuantity  int     `json:"totalBuyQuantity"`
	TotalSellQuantity int     `json:"totalSellQuantity"`
	AtoBuyQty         int     `json:"atoBuyQty"`
	AtoSellQty        int     `json:"atoSellQty"`
	Change            float64 `json:"Change"`
	PerChange         float64 `json:"perChange"`
	PrevClose         float64 `json:"prevClose"`
}

type EquityDetails struct {
	Info          EquityInfo          `json:"info"`
	Metadata      EquityMetadata      `json:"metadata"`
	SecurityInfo  EquitySecurityInfo  `json:"securityInfo"`
	PriceInfo     EquityPriceInfo     `json:"priceInfo"`
	PreOpenMarket EquityPreOpenMarket `json:"preOpenMarket"`
	SddDetails    SddDetails          `json:"sddDetails"`
	IndustryInfo  IndustryInfo        `json:"industryInfo"`
}

type SddDetails struct {
	SDDAuditor string `json:"SDDAuditor"`
	SDDStatus  string `json:"SDDStatus"`
}

type IndustryInfo struct {
	Macro         string `json:"macro"`
	Sector        string `json:"sector"`
	Industry      string `json:"industry"`
	BasicIndustry string `json:"basicIndustry"`
}

type EquityTradeInfo struct {
	NoBlockDeals   bool `json:"noBlockDeals"`
	BulkBlockDeals []struct {
		Name string `json:"name"`
	} `json:"bulkBlockDeals"`
	MarketDeptOrderBook struct {
		TotalBuyQuantity  int `json:"totalBuyQuantity"`
		TotalSellQuantity int `json:"totalSellQuantity"`
		Bid               []struct {
			Price    float64 `json:"price"`
			Quantity int     `json:"quantity"`
		} `json:"bid"`
		Ask []struct {
			Price    float64 `json:"price"`
			Quantity int     `json:"quantity"`
		} `json:"ask"`
		TradeInfo struct {
			TotalTradedVolume int     `json:"totalTradedVolume"`
			TotalTradedValue  float64 `json:"totalTradedValue"`
			TotalMarketCap    float64 `json:"totalMarketCap"`
			Ffmc              float64 `json:"ffmc"`
			ImpactCost        float64 `json:"impactCost"`
		} `json:"tradeInfo"`
		ValueAtRisk struct {
			SecurityVar       float64 `json:"securityVar"`
			IndexVar          float64 `json:"indexVar"`
			VarMargin         float64 `json:"varMargin"`
			ExtremeLossMargin float64 `json:"extremeLossMargin"`
			AdhocMargin       float64 `json:"adhocMargin"`
			ApplicableMargin  float64 `json:"applicableMargin"`
		} `json:"valueAtRisk"`
	} `json:"marketDeptOrderBook"`
	SecurityWiseDP struct {
		QuantityTraded           int    `json:"quantityTraded"`
		DeliveryQuantity         int    `json:"deliveryQuantity"`
		DeliveryToTradedQuantity int    `json:"deliveryToTradedQuantity"`
		SeriesRemarks            string `json:"seriesRemarks"`
		SecWiseDelPosDate        string `json:"secWiseDelPosDate"`
	} `json:"securityWiseDP"`
}

type DirectoryDetails struct {
	WebAddress string `json:"webAddress"`
	SMName     string `json:"smName"`
	Symbol     string `json:"symbol"`
	Office     string `json:"office"`
	Address    string `json:"address"`
	City       string `json:"city"`
	Pincode    string `json:"pincode"`
	Telephone  string `json:"telephone"`
	Fax        string `json:"fax"`
	Email      string `json:"email"`
}

type EquityCorporateInfo struct {
	Corporate struct {
		Announcements []struct {
			Desc         string `json:"desc"`
			AttchmntText string `json:"attchmntText"`
			AttchmntFile string `json:"attchmntFile"`
			AnDt         string `json:"an_dt"`
		} `json:"announcements"`
		BoardMeetings []struct {
			BmPurpose   string `json:"bm_purpose"`
			BmDesc      string `json:"bm_desc"`
			Attachment  string `json:"attachment"`
			BmDate      string `json:"bm_date"`
			BmTimestamp string `json:"bm_timestamp"`
		} `json:"boardMeetings"`
		CorporateActions []struct {
			Series      string `json:"series"`
			FaceVal     string `json:"faceVal"`
			Subject     string `json:"subject"`
			ExDate      string `json:"exDate"`
			RecDate     string `json:"recDate"`
			BcStartDate string `json:"bcStartDate"`
			BcEndDate   string `json:"bcEndDate"`
			NdStartDate string `json:"ndStartDate"`
			NdEndDate   string `json:"ndEndDate"`
		} `json:"corporateActions"`
		Governance           []interface{} `json:"governance"`
		FinancialResults     []interface{} `json:"financialResults"`
		ShareholdingPatterns struct {
			Cols []interface{} `json:"cols"`
			Data []interface{} `json:"data"`
		} `json:"shareholdingPatterns"`
		InsiderTrading          []interface{} `json:"insiderTrading"`
		SastRegulations29       []interface{} `json:"sastRegulations_29"`
		SastRegulations3132Post []interface{} `json:"sastRegulations_3132Post"`
		VotingResults           []interface{} `json:"votingResults"`
		AnnualReport            []struct {
			CompanyName string `json:"companyName"`
			FromYr      string `json:"fromYr"`
			ToYr        string `json:"toYr"`
			FileName    string `json:"fileName"`
		} `json:"annualReport"`
		DailyBuyBack        []interface{}      `json:"dailyBuyBack"`
		CompanyDirectory    []DirectoryDetails `json:"companyDirectory"`
		TransferAgentDetail []DirectoryDetails `json:"transferAgentDetail"`
		InvestorComplaints  []interface{}      `json:"investorComplaints"`
		Pledgedetails       []interface{}      `json:"pledgedetails"`
		CorpEncumbrance     []interface{}      `json:"corpEncumbrance"`
		SecretarialCamp     []interface{}      `json:"secretarialCamp"`
	} `json:"corporate"`
}

type EquityHistoricalInfo struct {
	ID               string  `json:"_id"`
	CHSymbol         string  `json:"CH_SYMBOL"`
	CHSeries         string  `json:"CH_SERIES"`
	CHMarketType     string  `json:"CH_MARKET_TYPE"`
	CHTradeHighPrice float64 `json:"CH_TRADE_HIGH_PRICE"`
	CHTradeLowPrice  float64 `json:"CH_TRADE_LOW_PRICE"`
	CHOpeningPrice   float64 `json:"CH_OPENING_PRICE"`
	CHClosingPrice   float64 `json:"CH_CLOSING_PRICE"`

	CHLastTradedPrice  float64 `json:"CH_LAST_TRADED_PRICE"`
	CHPreviousClsPrice float64 `json:"CH_PREVIOUS_CLS_PRICE"`
	CHTotTradedQty     float64 `json:"CH_TOT_TRADED_QTY"`
	CHTotTradedVal     float64 `json:"CH_TOT_TRADED_VAL"`
	CH52WeekHighPrice  float64 `json:"CH_52WEEK_HIGH_PRICE"`
	CH52WeekLowPrice   float64 `json:"CH_52WEEK_LOW_PRICE"`
	CHTotalTrades      float64 `json:"CH_TOTAL_TRADES"`
	CHISIN             string  `json:"CH_ISIN"`
	CHTimestamp        string  `json:"CH_TIMESTAMP"`
	Timestamp          string  `json:"TIMESTAMP"`
	CreatedAt          string  `json:"createdAt"`
	UpdatedAt          string  `json:"updatedAt"`
	VWAP               float64 `json:"VWAP"`
	MTimestamp         string  `json:"mTIMESTAMP"`
}

type EquityHistoricalData struct {
	Data []EquityHistoricalInfo `json:"data"`
	Meta struct {
		Series   []string `json:"series"`
		FromDate string   `json:"fromDate"`
		ToDate   string   `json:"toDate"`
		Symbols  []string `json:"symbols"`
	} `json:"meta"`
}

type IndexHistoricalData struct {
	Data struct {
		IndexCloseOnlineRecords []struct {
			EODCloseIndexVal float64 `json:"EOD_CLOSE_INDEX_VAL"`
			EODHighIndexVal  float64 `json:"EOD_HIGH_INDEX_VAL"`
			EODIndexName     string  `json:"EOD_INDEX_NAME"`
			EODLowIndexVal   float64 `json:"EOD_LOW_INDEX_VAL"`
			EODOpenIndexVal  float64 `json:"EOD_OPEN_INDEX_VAL"`
			EODTimestamp     string  `json:"EOD_TIMESTAMP"`
			TIMESTAMP        string  `json:"TIMESTAMP"`
		} `json:"indexCloseOnlineRecords"`
		IndexTurnoverRecords []struct {
			HITIndexNameUpper string  `json:"HIT_INDEX_NAME_UPPER"`
			HITTimestamp      string  `json:"HIT_TIMESTAMP"`
			HITTradedQty      float64 `json:"HIT_TRADED_QTY"`
			HITTurnOver       float64 `json:"HIT_TURN_OVER"`
			TIMESTAMP         string  `json:"TIMESTAMP"`
		} `json:"indexTurnoverRecords"`
	} `json:"data"`
}

type SeriesData struct {
	Data []string `json:"data"`
}

type IndexEquityInfo struct {
	Priority          int     `json:"priority"`
	Symbol            string  `json:"symbol"`
	Identifier        string  `json:"identifier"`
	Series            string  `json:"series"`
	Open              float64 `json:"open"`
	DayHigh           float64 `json:"dayHigh"`
	DayLow            float64 `json:"dayLow"`
	LastPrice         float64 `json:"lastPrice"`
	PreviousClose     float64 `json:"previousClose"`
	Change            float64 `json:"change"`
	PChange           float64 `json:"pChange"`
	TotalTradedVolume float64 `json:"totalTradedVolume"`
	TotalTradedValue  float64 `json:"totalTradedValue"`
	LastUpdateTime    string  `json:"lastUpdateTime"`
	YearHigh          float64 `json:"yearHigh"`
	FFMC              float64 `json:"ffmc"`
	YearLow           float64 `json:"yearLow"`
	NearWKH           float64 `json:"nearWKH"`
	NearWKL           float64 `json:"nearWKL"`
	PerChange365d     float64 `json:"perChange365d"`
	Date365dAgo       string  `json:"date365dAgo"`
	Chart365dPath     string  `json:"chart365dPath"`
	Date30dAgo        string  `json:"date30dAgo"`
	PerChange30d      float64 `json:"perChange30d"`
	Chart30dPath      string  `json:"chart30dPath"`
	ChartTodayPath    string  `json:"chartTodayPath"`
	Meta              struct {
		Symbol              string        `json:"symbol"`
		CompanyName         string        `json:"companyName"`
		Industry            string        `json:"industry"`
		ActiveSeries        []string      `json:"activeSeries"`
		DebtSeries          []interface{} `json:"debtSeries"`
		TempSuspendedSeries []interface{} `json:"tempSuspendedSeries"`
		IsFNOSec            bool          `json:"isFNOSec"`
		IsCASec             bool          `json:"isCASec"`
		IsSLBSec            bool          `json:"isSLBSec"`
		IsDebtSec           bool          `json:"isDebtSec"`
		IsSuspended         bool          `json:"isSuspended"`
		IsETFSec            bool          `json:"isETFSec"`
		IsDelisted          bool          `json:"isDelisted"`
		Isin                string        `json:"isin"`
	} `json:"meta"`
}

type IndexDetails struct {
	Name    string `json:"name"`
	Advance struct {
		Declines  string `json:"declines"`
		Advances  string `json:"advances"`
		Unchanged string `json:"unchanged"`
	} `json:"advance"`
	Timestamp string            `json:"timestamp"`
	Data      []IndexEquityInfo `json:"data"`
	Metadata  struct {
		IndexName         string  `json:"indexName"`
		Open              float64 `json:"open"`
		High              float64 `json:"high"`
		Low               float64 `json:"low"`
		PreviousClose     float64 `json:"previousClose"`
		Last              float64 `json:"last"`
		PercChange        float64 `json:"percChange"`
		Change            float64 `json:"change"`
		TimeVal           string  `json:"timeVal"`
		YearHigh          float64 `json:"yearHigh"`
		YearLow           float64 `json:"yearLow"`
		TotalTradedVolume float64 `json:"totalTradedVolume"`
		TotalTradedValue  float64 `json:"totalTradedValue"`
		FfmcSum           float64 `json:"ffmc_sum"`
	} `json:"metadata"`
	MarketStatus struct {
		Market              string  `json:"market"`
		MarketStatus        string  `json:"marketStatus"`
		TradeDate           string  `json:"tradeDate"`
		Index               string  `json:"index"`
		Last                float64 `json:"last"`
		Variation           float64 `json:"variation"`
		PercentChange       float64 `json:"percentChange"`
		MarketStatusMessage string  `json:"marketStatusMessage"`
	} `json:"marketStatus"`
	Date30dAgo  string `json:"date30dAgo"`
	Date365dAgo string `json:"date365dAgo"`
}
