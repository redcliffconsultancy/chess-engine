package main

import (
	"fmt"
	"strconv"
)

// 1 bit represents the piece
// location on the board
var bitBoard uint64

func setBit(n uint64, pos int) uint64 {
	// bitwise operators
	// | OR : Takes two numbers as operands and does OR on every bit of two numbers
	// << left shift: Takes two numbers
	n |= (1 << pos)
	return n
}

func clearBit(n uint64, pos int) uint64 {
	var mask uint64 = ^(1 << pos)
	n &= mask
	return n
}

func main() {
	// the board
	// 64 spaces - 1 to 64
	// or an 8x8 array
	bitBoard := setBit(bitBoard, 6)
	fmt.Println(strconv.FormatUint(bitBoard, 2))

	bitBoard = setBit(bitBoard, 7)
	fmt.Println(strconv.FormatUint(bitBoard, 2))

	bitBoard = clearBit(bitBoard, 6)
	fmt.Println(strconv.FormatUint(bitBoard, 2))
}

// all pieces take up one space in the array
// The king can move in 8 possible directions

type King struct {
	position     string
	moves        string
	validMoves   bool
	illegalMoves bool
	capture      bool
}
