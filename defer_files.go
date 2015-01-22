package main

import (
	"fmt"
	"os"
)

func main() {

	myFile := createFile("./boo.txt")
	defer closeFile(myFile)
	writeFile(myFile, "blah blah BOO!")

}

func createFile(path string) *os.File {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	fmt.Println("File created")
	return file
}

func writeFile(f *os.File, data string) {
	n, err := fmt.Fprintln(f, data)
	if err != nil {
		panic(err)
	}
	fmt.Println(n, "characters written to", f.Name())
}

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		panic(err)
	}
}
