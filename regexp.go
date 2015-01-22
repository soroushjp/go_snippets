//Regular expressions

package main

import (
	//"bytes"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	match, err := regexp.MatchString("p([a-z]+)ch", "peach")
	if err != nil {
		panic(err)
	}

	fmt.Println(match)

	//Start by compiling regular expression to get *Regexp instance
	r, err := regexp.Compile(`p([a-z]+)ch`)
	if err != nil {
		panic(err)
	}

	//Now we can run anything against our regexp
	//To get single leftmost match
	fmt.Println(r.FindString("peach punch"))
	fmt.Println(r.FindStringIndex("peach punch"))
	fmt.Println(r.FindStringSubmatch("peach punch"))
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	//To get all or n matches. -1 means get all
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	fmt.Println(r.FindAllStringIndex("peach punch", -1))
	fmt.Println(r.FindAllStringSubmatch("peach punch", -1))
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch", -1))

	//Above are string methods. May also be used as []byte search
	fmt.Println(r.Match([]byte("peach")))

	//Can also be used for Replace. MustCompile gives one return item with panic for error
	fmt.Println(regexp.MustCompile(`p([a-z]+)ch`).ReplaceAllString("peach punch pinch", "pooch"))

	//Can use Replace with function transformation of matched text
	fmt.Println(r.ReplaceAllStringFunc("peach pooch pinch", strings.ToUpper))

	//Print the regular expression itself
	fmt.Println(r.String())
}
