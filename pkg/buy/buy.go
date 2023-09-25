package buy

import "fmt"

func buyRatio(base float64, ratio float64) (l0, l1, l2 float64) {
	l0 = base * (1 + 1/ratio)
	l1 = (1 - level1) * l0
	l2 = (1 - level2) * l0
	return
}

func buyRatio50(base float64) {
	l0, l1, l2 := buyRatio(base, ratio50)
	printRes(l0, l1, l2, "ratio50")
}

func buyRatio20(base float64) {
	l0, l1, l2 := buyRatio(base, ratio20)
	printRes(l0, l1, l2, "ratio20")
}

func buyRatio10(base float64) {
	l0, l1, l2 := buyRatio(base, ratio10)
	printRes(l0, l1, l2, "ratio10")
}

func printRes(l0, l1, l2 float64, ratio string) {
	fmt.Printf("%s: %f, %f, %f\n", ratio, l0, l1, l2)
}

func Exec(base float64) {
	buyRatio50(base)
	buyRatio20(base)
	buyRatio10(base)
}
