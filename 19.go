package main

import "fmt"

func main() {
	ele1 := [5]int{1, 3, 4, 5, 8}
	ele2 := [5]int{1, 5, 8, 10, 13}
	fmt.Println(ele1)
	fmt.Println(ele2)
	soma := make([]int, len(ele1))
	for i := 0; i < len(ele1); i++ {
		soma[i] = ele1[i] + ele2[i]
	}
	fmt.Println("A soma deles é: ", soma)
}
