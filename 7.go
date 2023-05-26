package main

import (
	"fmt"
	"sort"
)

func main() {
	ele := []int{4, 5, 8, 7, 6}
	fmt.Println(ele)
	var n int
	fmt.Println("Informe um número: ")
	fmt.Scanln(&n)
	existe := false
	for _, eles := range ele {
		if eles == n {
			existe = true
			break
		}
	}
	if existe {
		fmt.Println("Já existe")
	} else {
		ele = append(ele, n)
		sort.Ints(ele)
		fmt.Println(ele)
	}
}
