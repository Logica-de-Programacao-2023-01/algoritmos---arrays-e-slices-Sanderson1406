package main

import "fmt"

func main() {
	ele := []string{"Oi", "Bom dia", "Boa tarde", "Tudo bem", "Boa noite", "Obrigado", "Tchau", "Adeus"}
	fmt.Println(ele)
	var v string
	fmt.Println("Informe um valor: ")
	fmt.Scan(&v)
	nv := []string{}
	for _, eles := range ele {
		if eles != v {
			nv = append(nv, eles)
		}
	}
	fmt.Println(nv)

}
