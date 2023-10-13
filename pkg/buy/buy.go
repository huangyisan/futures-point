package buy

import (
	"encoding/json"
	"fmt"
)

func buyRatio(price float64, ratio float64) (l0, l1, l2 float64) {
	l0 = price * (1 + 1/ratio)
	l1 = (1 - level1) * l0
	l2 = (1 - level2) * l0
	return
}

func buyRatio50(price float64) Ratio50 {
	l0, l1, l2 := buyRatio(price, ratio50)
	return Ratio50{l0, l1, l2}
}

// buyRation20
func buyRatio20(price float64) Ratio20 {
	l0, l1, l2 := buyRatio(price, ratio20)
	return Ratio20{l0, l1, l2}
}

// buyRatio10
func buyRatio10(price float64) Ratio10 {
	l0, l1, l2 := buyRatio(price, ratio10)
	return Ratio10{l0, l1, l2}
}

// buyRatio05
func buyRatio05(price float64) Ratio05 {
	l0, l1, l2 := buyRatio(price, ratio05)
	return Ratio05{l0, l1, l2}
}

// buyRatio02
func buyRatio02(price float64) Ratio02 {
	l0, l1, l2 := buyRatio(price, ratio02)
	return Ratio02{l0, l1, l2}
}

// buyRatio01
func buyRatio01(price float64) Ratio01 {
	l0, l1, l2 := buyRatio(price, ratio01)
	return Ratio01{l0, l1, l2}
}

// Exec
func Exec(price float64) []byte {
	buyRatio50(price)
	buyRatio20(price)
	buyRatio10(price)
	buyRatio05(price)
	buyRatio02(price)
	buyRatio01(price)

	j := JSONData{
		Side: "sell",
		Data: Data{buyRatio50(price), buyRatio20(price), buyRatio10(price), buyRatio05(price), buyRatio02(price), buyRatio01(price)},
	}
	jsonData, err := json.Marshal(j)
	if err != nil {
		fmt.Println(err)
		return jsonData
	}
	return jsonData
}
