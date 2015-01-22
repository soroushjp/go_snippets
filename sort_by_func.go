//Sorting by Functions

package main

import (
	"fmt"
	"sort"
)

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	words := []string{"boo", "bahss", "sdasdasd", "ab", "a", "aaaa"}

	sort.Sort(ByLength(words))

	fmt.Println(words)
}
