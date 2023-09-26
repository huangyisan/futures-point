package buy

import (
	"fmt"
	"testing"
)

func Test_buyPrice50(t *testing.T) {
	buyPrice50(100, 50)
}

func Test_buyPrice20(t *testing.T) {
	buyPrice20()
}

func buyPrice50(price float64, ratio float64) {
	l0 := price * (1 + 1/ratio)
	fmt.Println(l0)
}

func buyPrice20() {
	var prices float64 = 100
	var ratio float64 = 50
	l0 := prices * (1 + 1/ratio)
	fmt.Println(l0)
}
