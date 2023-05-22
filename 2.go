package main

import "fmt"

func main() {
	n := []int{1, 2, 3, 4, 5}
	fmt.Println(n)
	n = append(n[:2], n[3:]...)
	fmt.Println(n)
}
