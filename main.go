package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

func main() {
	// Heres a 120-bit sequence represented as a hexidecimal string
	// Let decode the hex string into bytes
	input := "1234567890abcdef1a2b3c4d5e6f70"
	inputArray, _ := hex.DecodeString(input)
	fmt.Println(len(inputArray))
	fmt.Println(inputArray)
	fmt.Println()

	// Lets create some registers based on the inputArray
	// all 80 bits in length
	wRegister := append(inputArray[0:5], inputArray[0:5]...)
	xRegister := inputArray[5:15]
	yRegister := make([]byte, 10)
	zRegister := append(inputArray[10:15], inputArray[0:5]...)
	fmt.Println(wRegister)
	fmt.Println(xRegister)
	fmt.Println(yRegister)
	fmt.Println(zRegister)
	fmt.Println()

	// What's the biggest number with can store as a uint64 data type
	number := uint64(18446744073709551615)
	fmt.Printf("Here's a big number: %d\n", number)
	fmt.Printf("Is this possible: %d\n", number-1)
	fmt.Printf("Is this possible: %d\n", number+1)

	// Let's represent this as 7 bytes. i.e. 7x8=56bits
	hiBytes := append([]byte{0, 0, 0}, inputArray[0:5]...)
	loBytes := append([]byte{0, 0, 0}, inputArray[5:10]...)
	loNumber := binary.BigEndian.Uint64(loBytes)
	hiNumber := binary.BigEndian.Uint64(hiBytes)

	// Let's print out some results

	fmt.Printf("\n%v \n", loBytes)
	fmt.Printf("%v\n", hiBytes)
	fmt.Printf("Here's the max 40 bit number: %d\n", loNumber)
	fmt.Printf("Here's the max 40 bit number: %d\n", hiNumber)
}
