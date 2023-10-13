package sell

import (
	"encoding/json"
	"fmt"
)

func sellRatio(price float64, ratio float64) (l0, l1, l2 float64) {
	l0 = price * (1 - 1/ratio)
	l1 = (1 + level1) * l0
	l2 = (1 + level2) * l0
	return
}

func sellRatio50(prices float64) Ratio50 {
	l0, l1, l2 := sellRatio(prices, ratio50)
	return Ratio50{l0, l1, l2}
}

func sellRatio20(price float64) Ratio20 {
	l0, l1, l2 := sellRatio(price, ratio20)
	return Ratio20{l0, l1, l2}
}

func sellRatio10(price float64) Ratio10 {
	l0, l1, l2 := sellRatio(price, ratio10)
	return Ratio10{l0, l1, l2}
}

func sellRatio05(price float64) Ratio05 {
	l0, l1, l2 := sellRatio(price, ratio05)
	return Ratio05{l0, l1, l2}
}

func sellRatio02(price float64) Ratio02 {
	l0, l1, l2 := sellRatio(price, ratio02)
	return Ratio02{l0, l1, l2}
}

func sellRatio01(price float64) Ratio01 {
	l0, l1, l2 := sellRatio(price, ratio01)
	return Ratio01{l0, l1, l2}
}

func printRes(l0, l1, l2 float64, ratio string) string {
	return fmt.Sprintf("%s: %f, %f, %f\n", ratio, l0, l1, l2)
}

func Exec(price float64) []byte {

	sellRatio50(price)
	sellRatio20(price)
	sellRatio10(price)
	sellRatio05(price)
	sellRatio02(price)
	sellRatio01(price)
	j := JSONData{
		Side: "buy",
		Data: Data{sellRatio50(price), sellRatio20(price), sellRatio10(price), sellRatio05(price), sellRatio02(price), sellRatio01(price)},
	}
	jsonData, err := json.Marshal(j)
	if err != nil {
		fmt.Println(err)
		return jsonData
	}
	return jsonData
}
