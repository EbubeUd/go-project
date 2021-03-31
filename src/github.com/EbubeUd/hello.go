package main

import (
	"fmt"
	"strconv"
)

var hello string = "local/private"
var Hello string = "Public"

func main() {
	var i float32 = 20.
	k := 9

	k = int(i)
	hello = strconv.Itoa(k)

	fmt.Println(i)
	fmt.Printf("%v is %T", i, i)
	fmt.Printf("%v is %T", k, k)
	fmt.Printf("%v is %T", hello, hello)
}
