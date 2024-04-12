package ascii6bit

import (
	"encoding/hex"
	"fmt"
	"math"
	"math/bits"
)

func main() {
	hexString := "553A6C"
	decodedHexString := DecodeHexString(hexString)
	fmt.Printf("%s\n", decodedHexString[:])

	data, err := hex.DecodeString(hexString)
	if err != nil {
		panic(err)
	}

	charsRev := createSTArray(data)
	var charsDecoded = make([]string, 24)
	for i := 0; i < len(charsRev); i++ {
		charsDecoded[i] = LookupRpMax(charsRev[i])
	}
	fmt.Printf("%s\n", charsDecoded[:])

}

// Returns a bit value from the provided data byte array,
// where bit 0 is the MSB of the first byte.
func getBitValue(data []byte, b int) bool {
	bytePosition := int(float32(b) / 8)
	bitPosition := int((float32(b)/8 - float32(bytePosition)) * 8)
	mask := bits.Reverse8(byte(math.Pow(2, float64(bitPosition))))
	return (data[bytePosition] & mask) != 0
}

// Create a new slice of 144 bits (or 18 bytes)
// The slice is created in two parts from the provided data
// Firstly taking the bits 0-111 (14 bytes)
// Secondly taking bits 116-149 (4 bytes)
func createSTArray(data []byte) []byte {
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
	return STArray
}
