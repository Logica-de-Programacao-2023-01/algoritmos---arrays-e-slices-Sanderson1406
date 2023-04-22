package main

import "fmt"

func main() {
	ele := []int{4, 3, 5, 7, 8, 9, 11, 22, 17, 21}
	fmt.Println(ele)
	min := ele[0]
	max := ele[0]
	for _, eles := range ele {
		if eles < min {
			min = eles // atualiza o menor número encontrado
		}
		if eles > max {
			max = eles
		}
	}
	fmt.Println("O menor número é: ", min)
	fmt.Println("O maior número é: ", max)

}
