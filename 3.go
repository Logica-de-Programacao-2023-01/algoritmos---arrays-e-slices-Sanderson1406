package main

import "fmt"

func main() {
	n := [4]float64{1.4, 8.4, 4.7, 9.4}
	var prod float64
	for _, v := range n {
		prod *= v
	}
	fmt.Println(prod)
}
