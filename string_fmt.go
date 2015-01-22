package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}

	fmt.Printf("%v\n", p)                           //%v any Go type
	fmt.Printf("%+v\n", p)                          //%+v include struct names
	fmt.Printf("%#v\n", p)                          //%#v Go syntax representation
	fmt.Printf("%T\n", p)                           //%T type of var
	fmt.Printf("%t\n", true)                        //bools
	fmt.Printf("%d\n", 55)                          //standard int
	fmt.Printf("%b\n", 55)                          //binary representation
	fmt.Printf("%c\n", 33)                          //character for given int
	fmt.Printf("%x\n", 16)                          //hex representation
	fmt.Printf("%f\n", 32.9)                        //standard float
	fmt.Printf("%e\n", 32.9)                        //scientific notation
	fmt.Printf("%E\n", 32.9)                        //scientific notation w/ capital E
	fmt.Printf("%s\n", "\"blah\"")                  //strings
	fmt.Printf("%q\n", "\"blah\"")                  //strings with double quotes
	fmt.Printf("%p\n", &p)                          //pointers
	fmt.Printf("|%6d|%6d|\n", 45, 3322)             //Fixed width
	fmt.Printf("|%6.2f|%6.2f|\n", 45.1, 22.55532)   //Fixed width.precision floats
	fmt.Printf("|%-6.2f|%-6.2f|\n", 45.1, 22.55532) //Fixed width.precision floats w/ left alignment
	fmt.Printf("|%6s|%6s|\n", "blah", "do")         //Fixed width strings
	fmt.Printf("|%-6s|%-6s|\n", "blah", "do")       //Fixed width strings w/ left alignment

	s := fmt.Sprintf("Print to a string") //Sprintf outputs to a string
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error") //Fprintf prints to a file, which can be os.Stdout os.Stdin os.Stderr also

}
