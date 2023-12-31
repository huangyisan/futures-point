package sell

import "fmt"

func sellRatio(price float64, ratio float64) (l0, l1, l2 float64) {
	l0 = price * (1 - 1/ratio)
	l1 = (1 + level1) * l0
	l2 = (1 + level2) * l0
	return
}

func sellRatio50(prices float64) {
	l0, l1, l2 := sellRatio(prices, ratio50)
	printRes(l0, l1, l2, "ratio50")
}

func sellRatio20(price float64) {
	l0, l1, l2 := sellRatio(price, ratio20)
	printRes(l0, l1, l2, "ratio20")
}

func sellRatio10(price float64) {
	l0, l1, l2 := sellRatio(price, ratio10)
	printRes(l0, l1, l2, "ratio10")
}

func sellRatio05(price float64) {
	l0, l1, l2 := sellRatio(price, ratio05)
	printRes(l0, l1, l2, "ratio05")
}

func sellRatio02(price float64) {
	l0, l1, l2 := sellRatio(price, ratio02)
	printRes(l0, l1, l2, "ratio02")
}

func sellRatio01(price float64) {
	l0, l1, l2 := sellRatio(price, ratio01)
	printRes(l0, l1, l2, "ratio01")
}

func printRes(l0, l1, l2 float64, ratio string) {
	fmt.Printf("%s: %f, %f, %f\n", ratio, l0, l1, l2)
}

func Exec(price float64) {
	sellRatio50(price)
	sellRatio20(price)
	sellRatio10(price)
	sellRatio05(price)
	sellRatio02(price)
	sellRatio01(price)
}
