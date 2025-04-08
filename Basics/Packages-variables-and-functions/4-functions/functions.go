package main

import "fmt"

func add(x int, y int) int {
	return x + y
}


func minus(x int, y int) int {
	return x - y
}

func main() {
	fmt.Println(add(42, 10))
	fmt.Println(minus(92, 22))
}