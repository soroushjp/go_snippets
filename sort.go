//Sorting

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	myString := "whatupworld"

	strSlice := strings.Split(myString, "")
	sort.Strings(strSlice)
	sortedString := strings.Join(strSlice, "")
	fmt.Println(sortedString)

	ints := []int{1, 5, 23, 3, 88, 4, 17}
	sort.Ints(ints)
	fmt.Println(ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted bool: ", s)
}
