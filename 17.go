package main

import "fmt"

func main() {
	ele := [10]int{1, 5, 8, 12, 15, 20, 33, 50, 45, 11}
	fmt.Println(ele)
	soma := 0
	for i := 1; i < len(ele); i += 2 {
		soma += ele[i]
	}
	fmt.Println(soma)
}
