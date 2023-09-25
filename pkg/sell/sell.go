package sell

import "fmt"

func sellRatio(base float64, ratio float64) (l0, l1, l2 float64) {
	l0 = base * (1 - 1/ratio)
	l1 = (1 + level1) * l0
	l2 = (1 + level2) * l0
	return
}

func sellRatio50(base float64) {
	l0, l1, l2 := sellRatio(base, ratio50)
	printRes(l0, l1, l2, "ratio50")
}

func sellRatio20(base float64) {
	l0, l1, l2 := sellRatio(base, ratio20)
	printRes(l0, l1, l2, "ratio20")
}

func sellRatio10(base float64) {
	l0, l1, l2 := sellRatio(base, ratio10)
	printRes(l0, l1, l2, "ratio10")
}

func printRes(l0, l1, l2 float64, ratio string) {
	fmt.Printf("%s: %f, %f, %f\n", ratio, l0, l1, l2)
}

func Exec(base float64) {
	sellRatio50(base)
	sellRatio20(base)
	sellRatio10(base)
}
