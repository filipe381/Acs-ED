package main

import (
	"fmt"
	"math"
)

func main() {
	var N, H, C, L int

	fmt.Scan(&N)
	fmt.Scan(&H, &C, &L)

	comprimentoRampa := float64(N) * math.Sqrt(float64(H*H+C*C))
	areaRampa := comprimentoRampa * float64(L)
	areaRampaM2 := areaRampa / 10000

	fmt.Printf("%.4f\n", areaRampaM2)
}
