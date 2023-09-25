package main

import (
	"fmt"
	"futures-point/pkg/buy"
	"futures-point/pkg/sell"
)

func main() {
	fmt.Println("sell:")
	sell.Exec(24)
	fmt.Println("buy:")
	buy.Exec(24)
}
