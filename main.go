package main

import (
	"encoding/json"
	"fmt"
	"go-nse/lib/nse"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "NSE",
	Short: "MyApp is a sample command-line application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("use 'help' to know the use")
	},
}

var symbolCmd = &cobra.Command{
	Use:   "symbol",
	Short: "Get Symbols",
	Run: func(cmd *cobra.Command, args []string) {
		data := nse.GetSymbols()
		for _, v := range data {
			fmt.Println(v)
		}
	},
}

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Greet someone",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(`Use Get Symbols to get all the symbols
nse symbols
`)
	},
}

var quoteEquityCmd = &cobra.Command{
	Use:   "quote-equity",
	Short: "Get Quote Equity",
	Run: func(cmd *cobra.Command, args []string) {
		symbol, _ := cmd.Flags().GetString("symbol")
		data, _ := nse.QuoteEquity(symbol)
		jsonData, _ := json.MarshalIndent(data, "", "  ")
		fmt.Printf("%s\n", string(jsonData))
	},
}

func init() {
	quoteEquityCmd.Flags().StringP("symbol", "s", "Guest", "Specify the name")

	rootCmd.AddCommand(helpCmd, symbolCmd, quoteEquityCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
