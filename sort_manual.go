//Manual sorting

package main

import (
	"fmt"
)

func Swap(input []int, i, j int) {
	input[i], input[j] = input[j], input[i]
}

func InsertionSort(input []int) {
	for i := 1; i < len(input); i++ {
		for j := i; j > 0 && input[j] < input[j-1]; j-- {
			Swap(input, j, j-1)
		}
	}
}

func SelectionSort(input []int) {
	var min int

	for i := 0; i < len(input); i++ {
		min = i
		for j := i + 1; j < len(input); j++ {
			if input[j] < input[min] {
				min = j
			}
		}
		Swap(input, min, i)
	}
}

func BubbleSort(input []int) {

	didSwap := 0
	for {
		for i := 1; i < len(input); i++ {
			if input[i] < input[i-1] {
				Swap(input, i, i-1)
				didSwap = 1
			}
		}
		if didSwap == 0 {
			break
		} else {
			didSwap = 0
		}
	}
}

func main() {
	input := []int{23, 5, 12, 53, 2, 3, 18, 36, 7, 15}

	InsertionSort(input)

	fmt.Println(input)
}
