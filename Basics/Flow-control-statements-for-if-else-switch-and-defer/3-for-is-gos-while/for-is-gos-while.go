package main

import "fmt"

func main() {
	sum := 1
	for sum < 50 {
		sum += 1
		fmt.Println(sum)
	}
}