package main

import (
	"fmt"
	s "strings"
)

func main() {
	fmt.Println("Contains: ", s.Contains("test", "es"))
	fmt.Println("Count: ", s.Count("test", "t"))
	fmt.Println("HasPrefix: ", s.HasPrefix("test", "te"))
	fmt.Println("HasSuffix:", s.HasSuffix("test", "st"))
	fmt.Println("Index:", s.Index("test", "e"))
	fmt.Println("Join:", s.Join([]string{"Hello", "World"}, " "))
	fmt.Println("Repeat:", s.Repeat("blah ", 6))
	fmt.Println("Replace:", s.Replace("test", "t", "z", 1))
	fmt.Println("Replace:", s.Replace("test", "t", "d", -1))
	fmt.Println("Split:", s.Split("test", "e"))
	fmt.Println("ToLower:", s.ToLower("TEST"))
	fmt.Println("ToUpper:", s.ToUpper("test"))
}
