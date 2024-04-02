package ascii6bit

import (
	"fmt"
	"strconv"
)

// Returns an indexed string array of decoded
// 6-bit ascii characters.
func DecodeHexString(hexString string) []string {

	// Determine the number of characters in the hexString
	numChars := len(hexString) * 4 / 6

	// Convert the hex string to an integer.
	integerValue, err := strconv.ParseInt(hexString, 16, 64)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Every 6-bits represents an ascii character
	// Copy every 6 bits to an indexed byte of an array.
	var Chars = make([]byte, numChars)
	for i := 0; i < 6; i++ {
		for j := 0; j < numChars; j++ {
			if hasBit(integerValue, uint(i+(j*6))) {
				Chars[j] = setBit(Chars[j], i)
			}
		}
	}

	// Reverse the byte order
	charsRev := reverseByteOrder(Chars[:])

	// Lookup each byte against the lookup table.
	var charsDecoded = make([]string, numChars)
	for i := 0; i < len(charsRev); i++ {
		charsDecoded[i] = lookupRpMax(charsRev[i])
	}
	return charsDecoded[:]

}

func setBit(n byte, pos int) byte {
	n |= (1 << pos)
	return n
}

func hasBit(n int64, pos uint) bool {
	val := n & (1 << pos)
	return (val > 0)
}

func reverseByteOrder(ba []byte) []byte {
	for i, j := 0, len(ba)-1; i < j; i, j = i+1, j-1 {
		ba[i], ba[j] = ba[j], ba[i]
	}
	return ba
}

func lookupRpMax(val uint8) string {
	rpMaxRegisters := map[uint8]string{
		0x00: "null",
		0x01: "A",
		0x02: "B",
		0x03: "C",
		0x04: "D",
		0x05: "E",
		0x06: "F",
		0x07: "G",
		0x08: "H",
		0x09: "I",
		0x0A: "J",
		0x0B: "K",
		0x0C: "L",
		0x0D: "M",
		0x0E: "N",
		0x0F: "O",
		0x10: "P",
		0x11: "Q",
		0x12: "R",
		0x13: "S",
		0x14: "T",
		0x15: "U",
		0x16: "V",
		0x17: "W",
		0x18: "X",
		0x19: "Y",
		0x1A: "Z",
		0x1B: "[",
		0x1C: "\\",
		0x1D: "]",
		0x1E: "^",
		0x1F: "_",
		0x20: " ",
		0x21: "!",
		0x22: "\"",
		0x23: "#",
		0x24: "$",
		0x25: "%",
		0x26: "&",
		0x27: "'",
		0x28: "(",
		0x29: ")",
		0x2A: "*",
		0x2B: "+",
		0x2C: ",",
		0x2D: "-",
		0x2E: ".",
		0x2F: "/",
		0x30: "0",
		0x31: "1",
		0x32: "2",
		0x33: "3",
		0x34: "4",
		0x35: "5",
		0x36: "6",
		0x37: "7",
		0x38: "8",
		0x39: "9",
		0x3A: ":",
		0x3B: ";",
		0x3C: "<",
		0x3D: "=",
		0x3E: ">",
		0x3F: "?",
	}
	return rpMaxRegisters[val]
}
