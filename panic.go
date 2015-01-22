package main

import (
	"fmt"
	"os"
)

func main() {

	//panic("oh oh something wen't wrong")

	file, err := os.Create("/path/to/something/that/doesnt/exist")
	if err != nil {
		panic(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
		fmt.Println("defer called")
	}()

}
