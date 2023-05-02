package main

import "fmt"

func main() {
	n := []int{1, 2, 3, 4, 5}
	fmt.Println(n)
	n = append(n[:3], n[4:]...)
	fmt.Println(n)
}
