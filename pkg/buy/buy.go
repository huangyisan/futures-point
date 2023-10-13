package buy

import "fmt"

func buyRatio(price float64, ratio float64) (l0, l1, l2 float64) {
	l0 = price * (1 + 1/ratio)
	l1 = (1 - level1) * l0
	l2 = (1 - level2) * l0
	return
}

func buyRatio50(price float64) string {
	l0, l1, l2 := buyRatio(price, ratio50)
	return printRes(l0, l1, l2, "ratio50")
}

func buyRatio20(price float64) string {
	l0, l1, l2 := buyRatio(price, ratio20)
	return printRes(l0, l1, l2, "ratio20")
}

func buyRatio10(price float64) string {
	l0, l1, l2 := buyRatio(price, ratio10)
	return printRes(l0, l1, l2, "ratio10")
}

func buyRatio05(price float64) string {
	l0, l1, l2 := buyRatio(price, ratio05)
	return printRes(l0, l1, l2, "ratio05")
}

func buyRatio02(price float64) string {
	l0, l1, l2 := buyRatio(price, ratio02)
	return printRes(l0, l1, l2, "ratio02")
}

func buyRatio01(price float64) string {
	l0, l1, l2 := buyRatio(price, ratio01)
	return printRes(l0, l1, l2, "ratio01")
}

func printRes(l0, l1, l2 float64, ratio string) string {
	res := fmt.Sprintf("%s: %f, %f, %f\n", ratio, l0, l1, l2)
	return res
}

func Exec(price float64) string {
	buyRatio50(price)
	buyRatio20(price)
	buyRatio10(price)
	buyRatio05(price)
	buyRatio02(price)
	buyRatio01(price)
	return fmt.Sprintf("%s%s%s%s%s%s%s", "sell\n", buyRatio50(price), buyRatio20(price), buyRatio10(price), buyRatio05(price), buyRatio02(price), buyRatio01(price))
}
