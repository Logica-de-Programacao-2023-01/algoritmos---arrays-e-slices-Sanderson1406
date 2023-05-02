package main

import "fmt"

func main() {
	ele := []float64{3.1, 3.3, 4.0, 1.5, 4.5, 6.5}
	var n float64
	fmt.Println("Informe um número que subistituirá todos os elemntos: ")
	fmt.Scan(&n)
	for i := 0; i < len(ele); i++ {
		ele[i] = n
		fmt.Println(ele)
	}
}
