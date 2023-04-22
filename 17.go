package main

import "fmt"

func main() {
	ele := [10]int{1, 5, 8, 12, 15, 20, 33, 50, 45, 11}
	fmt.Println(ele)
	soma := 0
	for _, eles := range ele {
		for {
			soma += eles
		}
	}
	fmt.Println(soma)
}
