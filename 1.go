package main

import "fmt"

func main() {
	var n = [3]int{5, 5, 7}
	soma := 0
	for _, n := range n {
		soma += n
	}
	fmt.Print(soma)
}
