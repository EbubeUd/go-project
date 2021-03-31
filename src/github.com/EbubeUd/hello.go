package main

import (
	"fmt"
)

func main() {
	a := [5]uint8{1, 2, 3, 4, 5}
	b := &a
	c := a[:]
	d := a[:2]
	e := a[2:]
	f := a[2:4]
	g := a[:]
	g = append(g, []uint8{1, 2, 3, 4, 5, 6}...)
	h := len(g)
	i := make([]uint8, 10, 100)
	j := len(i)

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", e)
	fmt.Printf("%v\n", f)
	fmt.Printf("%v\n", g)
	fmt.Printf("%v\n", h)
	fmt.Printf("%v\n", i)
	fmt.Printf("%v\n", j)

}
