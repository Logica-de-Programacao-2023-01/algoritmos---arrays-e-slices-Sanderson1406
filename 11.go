package main

import "fmt"

func main() {
	matriz := [2][3]int{{2, 4, 7}, {5, 6, 8}}
	fmt.Println(matriz)
	var l, c int
	fmt.Println("Informe a linha do elemento: ")
	fmt.Scanln(&l)
	fmt.Println("Informe a coluna do elemento: ")
	fmt.Scanln(&c)
	for linha := range matriz {
		for coluna, elemento := range matriz[linha] {
			fmt.Println("Elenento da linha ", linha, "coluna ", coluna, "é: ", elemento)
		}
	}
}
