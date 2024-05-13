package main

import (
	"fmt"
)

func main() {
	var H1, M1, H2, M2 int
	for {
		fmt.Scan(&H1, &M1, &H2, &M2)
		if H1 == 0 && M1 == 0 && H2 == 0 && M2 == 0 {
			break
		}

		minutosIniciais := H1*60 + M1
		minutosFinais := H2*60 + M2

		if minutosFinais < minutosIniciais {
			minutosFinais += 24 * 60
		}

		fmt.Println(minutosFinais - minutosIniciais)
	}
}
