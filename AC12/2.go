package main

import "fmt"

func main() {
	var rank int

	fmt.Scan(&rank)

	if rank <= 1 {
		fmt.Println("Top 1")
	} else if rank <= 3 {
		fmt.Println("Top 3")
	} else if rank <= 5 {
		fmt.Println("Top 5")
	} else if rank <= 10 {
		fmt.Println("Top 10")
	} else if rank <= 25 {
		fmt.Println("Top 25")
	} else if rank <= 50 {
		fmt.Println("Top 50")
	} else if rank <= 100 {
		fmt.Println("Top 100")
	} else {
		fmt.Println("Rank está fora do alcance")
	}
}
