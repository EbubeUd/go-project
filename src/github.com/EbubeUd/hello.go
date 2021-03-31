package main

import "fmt"

func main() {

	if false {
		fmt.Println("Hello")
	}

	switch 6 {
	case 1, 2, 3, 4, 5:
		fmt.Println("true")
	default:
		fmt.Println("false")
	}

}
