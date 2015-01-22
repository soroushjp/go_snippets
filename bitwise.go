package main

import "fmt"

func getBit(input int8, bit uint) int {

	if bit < 0 || bit > 7 {
		panic("Requested bit must be between 0 and 7")
	}

	return Btoi(input&(1<<bit) > 0)

}

func getBits(input int8) string {
	var bit int
	var i uint
	var bitString string = ""
	for i = 0; i < 8; i++ {
		bit = getBit(input, 7-i)
		bitString += fmt.Sprintf("%d", bit)
	}

	return bitString
}

func setBit(input int, bit uint) int {
	return input | (1 << bit)
}

func unsetBit(input int, bit uint) int {
	return input & ^(1 << bit)
}

func flipBit(input int, bit uint) int {
	return input ^ (1 << bit)
}

func Btoi(testBool bool) int {
	if testBool {
		return 1
	} else {
		return 0
	}
}

func main() {
	var a int = 12
	a = unsetBit(a, 3)

	fmt.Println(a)

}
