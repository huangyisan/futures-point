package buy

const (
	level0 float64 = 0
	level1 float64 = 0.0025
	level2 float64 = 0.005

	ratio50 = 50
	ratio20 = 20
	ratio10 = 10
	ratio05 = 5
	ratio02 = 2
	ratio01 = -1
)

type Ratio50 []float64
type Ratio20 []float64
type Ratio10 []float64
type Ratio05 []float64
type Ratio02 []float64
type Ratio01 []float64

type Data struct {
	Ratio50 `json:"ratio50"`
	Ratio20 `json:"ratio20"`
	Ratio10 `json:"ratio10"`
	Ratio05 `json:"ratio05"`
	Ratio02 `json:"ratio02"`
	Ratio01 `json:"ratio01"`
}

type JSONData struct {
	Side string `json:"side"`
	Data Data   `json:"data"`
}
