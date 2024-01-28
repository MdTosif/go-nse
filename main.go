package main

import (
	"encoding/json"
	"fmt"
	"go-nse/lib/nse"
	"log"

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
  -s, --symbol string    Specify the symbol (default "Guest")

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
		jsonData, _ := json.MarshalIndent(data, "", "  ")
		fmt.Printf("%s\n", string(jsonData))
	},
}

func init() {
	quoteEquityCmd.Flags().StringP(symbolFlagName, symbolFlagShort, symbolFlagDefault, symbolFlagDescription)

	rootCmd.AddCommand(helpCmd, symbolCmd, quoteEquityCmd)
	nse.Lol()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
