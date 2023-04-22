package main

import "fmt"

func main() {
	ele := []string{"Oi", "Bom dia", "Boa tarde", "Boa tarde", "Boa noite", "Obrigado", "Tchau", "Adeus"}
	fmt.Println(ele)
	var v string
	fmt.Println("Informe um valor: ")
	fmt.Scan(&v)
	for i := 0; i < len(ele); i++ {
		if ele[i] == v {
			ele = append(ele[:i], ele[i+1])
		}
	}
	fmt.Println(ele)
}
