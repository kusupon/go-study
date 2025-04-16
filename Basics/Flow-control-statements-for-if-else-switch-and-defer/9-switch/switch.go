package main 

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on")
	switch os := runtime.GOOS; os {
	case "drawin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("lunux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}