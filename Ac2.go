package main

import (
	"fmt"
)

func main() {
	vetor := [10]int{}
	for i := 0; i < 10; i++ {
		fmt.Printf("Digite o %dº número: ", i+1)
		fmt.Scanln(&vetor[i])
	}

	for i := 0; i < len(vetor); i++ {
		fmt.Println("O valor na posição ", i, " é:", vetor[i])
	}

}
