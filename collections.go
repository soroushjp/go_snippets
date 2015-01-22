//Collection functions

package main

import (
	"fmt"
	"strings"
)

func Index(vs []string, s string) int {
	for i, e := range vs {
		if e == s {
			return i
		}
	}
	return -1
}

func Include(vs []string, s string) bool {
	return Index(vs, s) >= 0
}

func Any(vs []string, f func(string) bool) bool {
	for _, e := range vs {
		if f(e) {
			return true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool {
	for _, e := range vs {
		if !f(e) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	filtered := []string{}
	for _, e := range vs {
		if f(e) {
			filtered = append(filtered, e)
		}
	}
	return filtered
}

func Map(vs []string, f func(string) string) []string {
	mapped := make([]string, len(vs))
	for i, e := range vs {
		mapped[i] = f(e)
	}
	return mapped
}

func main() {
	testSlice := []string{"boo", "bah", "gahgah", "booboo"}

	fmt.Println("testSlice:", testSlice)

	fmt.Println("Does slice include 'boo' ? :", Include(testSlice, "boo"))

	fmt.Println("Where does it include 'boo' ? :", Index(testSlice, "boo"))

	fmt.Println("Do any of the strings start with 'a' ? :", Any(testSlice, func(s string) bool {
		return string(s[0]) == "a"
	}))
	fmt.Println("Do any of the strings start with 'g' ? :", Any(testSlice, func(s string) bool {
		return string(s[0]) == "g"
	}))

	fmt.Println("Do all the strings contain the letter b? : ", All(testSlice, func(s string) bool {
		for _, e := range s {
			if string(e) == "b" {
				return true
			}
		}
		return false
	}))
	fmt.Println("Do all of the strings have 3 or more characters? :", All(testSlice, func(s string) bool {
		return len(s) >= 3
	}))

	fmt.Println("Filter out any string without the substring 'boo' in it somewhere", Filter(testSlice, func(s string) bool {
		substring := "boo"
		for i := 0; i <= (len(s) - len(substring)); i++ {
			match := 1
			for j := 0; j < len(substring); j++ {
				if s[i+j] != substring[j] {
					match = 0
					break
				}
			}
			if match == 1 {
				return true
			}
		}
		return false
	}))

	fmt.Println("Filter out any string without the substring 'ah' in it somewhere", Filter(testSlice, func(s string) bool {
		return strings.Contains(s, "ah")
	}))

	fmt.Println("Map by concatenating each string to itself: ", Map(testSlice, func(s string) string {
		return strings.Join([]string{s, s}, "")
	}))
}
