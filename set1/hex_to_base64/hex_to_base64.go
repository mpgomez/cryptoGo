package main

import (
	b64 "encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	// First, we declare our hex as a string
	hexStr := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	// Now, we need to convert it to actual binary hex
	hexB, err := hex.DecodeString(hexStr)
	if err != nil {
		panic(err)
	}
	// We can print the hex
	fmt.Printf("THE HEX: % x\n", hexB)
	// And also the string it represents,
	fmt.Printf("THE STRING: % s\n", hexB)

	// Now we base64 encode.

	sEnc := b64.StdEncoding.EncodeToString(hexB)
	fmt.Printf("THE BASE 64 AS STRING:  % s\n", sEnc)

	// Note: we could use func (enc *Encoding) Encode(dst, src []byte) instead.
	//var bEnc []byte
	bEnc := make([]byte, 64)
	b64.StdEncoding.Encode(bEnc, hexB)
	fmt.Printf("THE BASE 64 FROM BYTES: % s\n", bEnc)
}
