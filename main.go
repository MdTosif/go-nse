package main

import (
	"fmt"
	"log"
	"nse/lib/nse"

	"github.com/spf13/cobra"
)

const (
	rootCmdUse            = "NSE"
	rootCmdShort          = "MyApp is a sample command-line application"
	symbolCmdUse          = "symbol"
	symbolCmdShort        = "Get Symbols"
	helpCmdUse            = "help"
	helpCmdShort          = "Greet someone"
	quoteEquityCmdUse     = "quote-equity"
	quoteEquityCmdShort   = "Get Quote Equity"
	symbolFlagName        = "symbol"
	symbolFlagShort       = "s"
	symbolFlagDefault     = "Guest"
	symbolFlagDescription = "Specify the symbol"
)

var rootCmd = &cobra.Command{
	Use:   rootCmdUse,
	Short: rootCmdShort,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use 'help' to know the use")
	},
}

var symbolCmd = &cobra.Command{
	Use:   symbolCmdUse,
	Short: symbolCmdShort,
	Run: func(cmd *cobra.Command, args []string) {
		data := nse.GetSymbols()
		for _, v := range data {
			fmt.Println(v)
		}
	},
}

var helpCmd = &cobra.Command{
	Use:   helpCmdUse,
	Short: helpCmdShort,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`Usage:
  nse symbol          Get all symbols
  nse quote-equity    Get Quote Equity for a symbol

Flags:
  -s, --symbol string    Specify the symbol

Examples:
  nse symbol
  nse quote-equity --symbol TATATECH`)
	},
}

var quoteEquityCmd = &cobra.Command{
	Use:   quoteEquityCmdUse,
	Short: quoteEquityCmdShort,
	Run: func(cmd *cobra.Command, args []string) {
		symbol, _ := cmd.Flags().GetString(symbolFlagName)
		data, _ := nse.QuoteEquity(symbol)
		fmt.Printf("Company: %s (%s)\n", data.Info.CompanyName, data.Info.Symbol)
		fmt.Printf("Industry: %s\n", data.Info.Industry)
		fmt.Printf("Listing Date: %s\n", data.Metadata.ListingDate)
		fmt.Printf("Last Price: ₹%.2f\n", data.PriceInfo.LastPrice)
		fmt.Printf("Change: +%.2f (%.2f%%)\n", data.PreOpenMarket.Change, data.PriceInfo.PChange)
		fmt.Printf("Trading Status: %s\n", data.SecurityInfo.TradingStatus)
		fmt.Printf("Total Traded Volume: %d\n", data.PreOpenMarket.TotalTradedVolume)
		fmt.Printf("Trading Segment: %s\n", data.SecurityInfo.TradingSegment)
		fmt.Printf("Face Value: ₹%.2f\n", data.SecurityInfo.FaceValue)
		fmt.Printf("Issued Size: %.2f\n", data.SecurityInfo.IssuedSize)
		fmt.Printf("Week High: ₹%.2f\n", data.PriceInfo.WeekHighLow.Max)
		fmt.Printf("Week Low: ₹%.2f\n", data.PriceInfo.WeekHighLow.Min)
	},
}

func init() {
	quoteEquityCmd.Flags().StringP(symbolFlagName, symbolFlagShort, symbolFlagDefault, symbolFlagDescription)

	rootCmd.AddCommand(helpCmd, symbolCmd, quoteEquityCmd)

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
