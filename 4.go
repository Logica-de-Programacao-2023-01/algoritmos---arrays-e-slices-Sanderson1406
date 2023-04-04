package main

import "fmt"

func main() {
	var slice []int
	var soma, tamanho, x int
	fmt.Println("digite o tamanho do slice: ")
	fmt.Scanln(&tamanho)
	for i := 0; i < tamanho; i++ {
		fmt.Println("Digite um número ")
		fmt.Scanln(&x)
		slice = append(slice, x)
		soma = soma + x
	}
	fmt.Println(slice)
	fmt.Println(soma)
}
