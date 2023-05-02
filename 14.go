package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var x, y int
	fmt.Println("Digite o primeiro elemento ")
	fmt.Scan(&x)
	fmt.Println("Sgundo elemnto")
	fmt.Scan(&y)
	slice[x], slice[y] = slice[y], slice[x]
	fmt.Println(slice)
}
