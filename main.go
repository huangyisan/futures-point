//package main
//
//import (
//	"fmt"
//	"futures-point/pkg/buy"
//	"futures-point/pkg/sell"
//	"github.com/spf13/pflag"
//)
//
//var (
//	dir   string
//	price float64
//)
//
//func init() {
//	pflag.StringVarP(&dir, "dir", "d", "sell", "buy or sell")
//	pflag.Float64VarP(&price, "price", "p", 24, "price")
//
//	pflag.Parse()
//}
//func main() {
//	switch dir {
//	// sell is buy
//	case "sell":
//		fmt.Printf("sell: %f\n", price)
//		buy.Exec(price)
//		// buy is sell
//	case "buy":
//		fmt.Printf("buy: %f\n", price)
//		sell.Exec(price)
//	}
//}

package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "futures-point",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
