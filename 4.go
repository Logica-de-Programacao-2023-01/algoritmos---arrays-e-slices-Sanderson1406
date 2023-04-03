package main

import "fmt"

func main() {
	n := []int{0}
	var t int = 0
	var v1 int = 0
	var v2 int = 0
	var v3 int = 0
	var v4 int = 0
	var v5 int = 0
	var v6 int = 0
	fmt.Println(n)
	fmt.Println("Qual o tamanho da Slice? ")
	fmt.Scanln(&t)
	fmt.Println("Digite os valor 1: ")
	fmt.Scanln(&v1)
	fmt.Println("Digite os valor 2: ")
	fmt.Scanln(&v2)
	fmt.Println("Digite os valor 3: ")
	fmt.Scanln(&v3)
	fmt.Println("Digite os valor 4: ")
	fmt.Scanln(&v4)
	fmt.Println("Digite os valor 5: ")
	fmt.Scanln(&v5)
	fmt.Println("Digite os valor 6: ")
	fmt.Scanln(&v6)
	n = append(n, v1, v2, v3, v4, v5, v6)
	fmt.Println(n)
}
