package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

type ID []byte

func hash(key string) ID {
	h := sha256.Sum256([]byte(key))
	return h[:]
}

func xor(a, b ID) []byte {
	c := make([]byte, len(a))
	for i := 0; i < len(a); i++ {
		c[i] = a[i] ^ b[i]
	}
	return c
}

func zeroPrefixLen(id []byte) int {
	for i := 0; i < len(id); i++ {
		for j := 0; j < 8; j++ {
			if (id[i]>>uint8(7-j))&0x1 != 0 {
				return i*8 + j
			}
		}
	}
	return len(id) * 8
}

func toBinary(data byte) string {
	bin := ""
	for i := 7; i >= 0; i-- {
		s := strconv.Itoa(int(data >> uint8(i) & 1))
		bin += s
	}
	return bin
}

func main() {
	hA := hash("abc")
	hB := hash("efg")

	fmt.Println(hA)
	fmt.Println(hB)

	xorr := xor(hA, hB)

	fmt.Println(xorr)
	fmt.Println(toBinary(xorr[0]))
	fmt.Println(zeroPrefixLen(xorr))
}
