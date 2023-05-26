package main

import "fmt"

func main() {
	ele := []string{"Oi", "Bom dia", "Boa tarde", "Boa tarde", "Boa noite", "Obrigado", "Tchau", "Adeus"}
	fmt.Println(ele)
	var v string
	fmt.Println("Informe um valor: ")
	fmt.Scan(&v)
	saida := make([]string, 8)
	for _, x := range ele {
		if x != v {
			saida = append(saida, x)
		} else if x == v {
			continue
		}
	}
	fmt.Println(saida)
}
