package main

import (
	"fmt"

	"github.com/yenoxx/voltkeep/encoder"
)

func main() {
	a := []byte{123, 228, 13, 37, 14, 88}

	e := encoder.CreateEncoder()
	aEnc := e.Begin("aboba", a).Encode().ChangePass("123").Encode().Bytes()
	aDec := e.Begin("123", aEnc).Decode().ChangePass("aboba").Decode().Bytes()

	fmt.Printf("%v\n", aEnc)
	fmt.Printf("%v\n", aDec)
}
