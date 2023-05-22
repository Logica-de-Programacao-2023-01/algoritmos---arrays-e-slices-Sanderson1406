package main

import "fmt"

func main() {
	var n = [3]int{5, 5, 5}
	soma := 0
	for _, ns := range n {
		soma += ns
	}
	fmt.Print(soma)
}
