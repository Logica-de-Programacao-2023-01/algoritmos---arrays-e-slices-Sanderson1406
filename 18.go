package main

import "fmt"

func main() {
	var n int
	fmt.Println("Digite o valor de n: ")
	fmt.Scan(&n)
	var primos []int
	for i := 2; i <= n; i++ {
		primo := true
		for x := 2; x < i; i++ {
			if i%x == 0 {
				primo = false
				break
			}
		}
		if primo {
			primos = append(primos, i)
		}
	}
	fmt.Println(primos)
}
