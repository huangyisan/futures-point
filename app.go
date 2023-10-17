package main

import (
	"context"
	"fmt"
	"futures-point/pkg/buy"
	"futures-point/pkg/sell"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Calculate(price float64) (res [][]byte) {

	buyRes := buy.Exec(price)
	sellRes := sell.Exec(price)
	res = append(res, buyRes, sellRes)
	fmt.Println(string(buyRes))
	fmt.Println(string(sellRes))

	return
}
