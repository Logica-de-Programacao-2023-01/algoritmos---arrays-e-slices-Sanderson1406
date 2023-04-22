package main

import "fmt"

func main() {
	ele := [5]int{0, 3, 6, 7, 9}
	nv := []int{}
	for _, eles := range ele {
		if eles%3 == 0 {
			nv = append(nv, eles)
		}
	}
	fmt.Println(ele)
	fmt.Println(nv)
}
