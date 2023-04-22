package main

import "fmt"

func main() {
	slice := []int{10, 9, 8, 7, 6, 5, 4, 3}
	crescente := true
	for i := 1; i < len(slice); i++ {
		atual := slice[i]
		anterior := slice[i-1]
		if anterior >= atual {
			crescente = false
			break
		}
	}
	if crescente {
		fmt.Println("Ele está em ordem crescente")
	} else {
		fmt.Println("Não está em ordem crescente")
	}
}
