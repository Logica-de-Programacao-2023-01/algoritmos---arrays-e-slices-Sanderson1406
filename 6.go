package main

import "fmt"

func main() {
	var ele = [10]int{1, 5, 9, 7, 5, 3, 4, 11, 22, 10}
	var v int
	fmt.Println("Digite um valor: ")
	fmt.Scanln(&v)
	var encontrado bool = true
	fmt.Println(encontrado)
	for _, eles := range ele {
		if eles == v {
			encontrado = true
			fmt.Println("O valor existe na Array")
			break
		} else if eles != v {
			fmt.Println("O valor não existe")
			encontrado = false
			break
		}

	}

}
