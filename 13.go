package main

import "fmt"

func main() {
	ele := [7]int{2, 3, 5, 9, 12, 14, 17}
	var n1, n2 int
	fmt.Println(ele)
	fmt.Println("Informe o número que será adicionado no primeiro: ")
	fmt.Scanln(&n1)
	fmt.Println("Informe o número que será adicionado no ultimo: ")
	fmt.Scanln(&n2)
	var l1 = []int{}
	var e1 = []int{n1}
	var e2 = []int{n2}
	for _, eles := range ele {
		eles++
		l2 := make([]int, eles)
		l1 = append(l1, e1...)
		l1 = append(l1, l2...)
		l1 = append(l1, e2...)
	}
	fmt.Println(l1)
}
