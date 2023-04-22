package main

import "fmt"

func main() {
	ele := []int{4, 5, 8, 7, 6}
	var n int
	fmt.Println("Informe um número: ")
	fmt.Scanln(&n)
	for _, eles := range ele {
		if eles != n {
			ele = append(ele, n)
			fmt.Println("ele não está presente, com o elemento, o slice fica assim: ", ele)
			break
		} else {
			fmt.Println("O elemento já existe")
			break
		}
	}
}
