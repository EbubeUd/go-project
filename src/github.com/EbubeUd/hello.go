package main

import (
	"fmt"
)

func main() {
	var n complex64 = 1 + 2i
	var j complex64 = complex(1, 2)
	var r rune = 'a'

	fmt.Printf("%v and %T\n", real(n), n)
	fmt.Printf("%v and %T\n", imag(j), j)
	fmt.Printf("%v and %T\n", r, r)
}
