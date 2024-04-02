package main

import (
	"encoding/hex"
	"example/ascii6bit"
	"fmt"
	"math"
	"math/bits"
)

func main() {
	hexString := "553A6C"
	decodedHexString := ascii6bit.DecodeHexString(hexString)
	fmt.Printf("%s\n", decodedHexString[:])

	data, err := hex.DecodeString(hexString)
	if err != nil {
		panic(err)
	}

	createSTArray(data)
}

// returns the value of bit position b from the data
func getBitValue(data []byte, b int) bool {
	bytePosition := int(float32(b) / 8)
	bitPosition := int((float32(b)/8 - float32(bytePosition)) * 8)
	mask := bits.Reverse8(byte(math.Pow(2, float64(bitPosition))))
	return (data[bytePosition] & mask) != 0
}

func createSTArray(data []byte) {
	// Create a new slice of 144 bits (or 24 bytes)
	var count float64 = 0
	var index int
	var STArray = make([]byte, 24)
	for i := 0; i < 24; i++ {
		index = i / 6
		bv := getBitValue(data, i)
		mask := bits.Reverse8(byte(math.Pow(2, count)))
		if bv {
			STArray[index] = STArray[index] | byte(mask)>>2
		}
		count = count + 1
		if count == 6 {
			count = 0
		}

	}
	fmt.Println(STArray)
}
