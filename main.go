package main

import (
	"fmt"

	"github.com/yenoxx/voltkeep/encoder"
)

func main() {
	a := []byte{123, 228, 13, 37, 14, 88}

	e := encoder.CreateEncoder()
	aEnc := e.Begin(1234567, a).Encode().ChangeKey(321).Encode().Bytes()
	aDec := e.Begin(321, aEnc).Decode().ChangeKey(1234567).Decode().Bytes()

	fmt.Printf("%v\n", aEnc)
	fmt.Printf("%v\n", aDec)
}
