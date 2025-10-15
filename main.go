package main

import (
	"fmt"
	"yxenc/encoder"
)

func main() {
	a := []byte{123, 228, 13, 37, 14, 88}

	e := encoder.CreateEncoder()
	aEnc := e.Begin(1234567, a).Encode().Bytes()
	aDec := e.Begin(1234567, aEnc).Decode().Bytes()

	fmt.Printf("%v\n", aEnc)
	fmt.Printf("%v\n", aDec)
}
