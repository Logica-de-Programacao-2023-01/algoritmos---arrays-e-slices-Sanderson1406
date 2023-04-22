package main

import "fmt"

func main() {
	ele := [5]int{7, 5, 6, 17, 21}
	nv := []int{}
	for _, eles := range ele {
		if eles%3 == 0 {
			nv = append(nv, eles)
		}
	}
	fmt.Println(ele)
	fmt.Println(nv)
}
