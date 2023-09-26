package main

import (
	"fmt"
	"futures-point/pkg/buy"
	"futures-point/pkg/sell"
	"github.com/spf13/pflag"
)

var (
	dir   string
	price float64
)

func init() {
	pflag.StringVarP(&dir, "dir", "d", "sell", "buy or sell")
	pflag.Float64VarP(&price, "price", "p", 24, "price")

	pflag.Parse()
}
func main() {
	switch dir {
	// sell is buy
	case "sell":
		fmt.Printf("sell: %f\n", price)
		buy.Exec(price)
		// buy is sell
	case "buy":
		fmt.Printf("buy: %f\n", price)
		sell.Exec(price)
	}
}
