package main

import "fmt"

func main() {
	var ele = [10]int{1, 5, 9, 7, 5, 3, 4, 11, 22, 10}
	fmt.Println(ele)
	var v int
	fmt.Println("Digite um valor: ")
	fmt.Scanln(&v)
	for i := 0; i < len(ele); i++ {
		if ele[i] == v {
			fmt.Println("O valor existe na Array")
			return
		}
	}
	fmt.Println("O valor não existe")
}
