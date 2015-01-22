//Time formatting

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	//Formatting according to RFC3339
	fmt.Println(t.Format(time.RFC3339))

	//RFC3339 to time.Time
	t1, err := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	if err != nil {
		panic(err)
	}
	fmt.Println(t1)

	//Formatting by giving example layout. Example must use this exact reference time Jan 2, 2006 MST 3:04PM
	fmt.Println(t.Format("Jan 2, 2006 at MST 15:04"))
}
