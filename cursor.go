package main

import (
	"fmt"
)

type node struct {
	key   []byte
	val   []byte
	left  *node
	right *node
}

func main() {
	a := []int{1, 2, 3}
	a = a[:0]
	fmt.Println(a)

	code := 0x0011
	fmt.Println(code)
}
