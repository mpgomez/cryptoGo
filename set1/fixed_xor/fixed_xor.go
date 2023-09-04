package main

import (
	"encoding/hex"
	"errors"
	"fmt"
)

func xorHex(arg1, arg2 []byte) ([]byte, error) {
	if len(arg1) != len(arg2) {
		return nil, errors.New("Not the same length")
	}
	xorHex := make([]byte, len(arg1))
	for i:= range xorHex {
		xorHex[i] = arg1[i]^arg2[i]
	}
	return xorHex, nil
}

func main() {
	// First, we declare our hex as a string
	hexStr1 := "1c0111001f010100061a024b53535009181c"
	hexStr2 := "686974207468652062756c6c277320657965"
	// Now, we need to convert it to actual binary hex
	hexB1, err := hex.DecodeString(hexStr1)
	if err != nil {
		panic(err)
	}
	hexB2, err := hex.DecodeString(hexStr2)
	if err != nil {
		panic(err)
	}
	xorB, err := xorHex(hexB1, hexB2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("THE HEX: % x\n", xorB)
	fmt.Println("THE HEX SHOULD BE: 746865206b696420646f6e277420706c6179")
	fmt.Printf("String A: %s \n", hexB1)
	fmt.Printf("String B: %s \n", hexB2)
	fmt.Printf("XOR STRING (for the lols): % s\n", xorB)
}