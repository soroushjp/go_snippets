//Time
package main

import (
	"fmt"
	"time"
)

func main() {

	//Get current time
	fmt.Println(time.Now())

	//Get time.Time object for any arbitrary time
	then := time.Date(2015, 1, 22, 10, 38, 0, 0, time.UTC)

	//Different representations of time.Time
	fmt.Println(then)
	fmt.Println(then.Year())
	fmt.Println(then.Month())
	fmt.Println(then.Day())
	fmt.Println(then.Weekday())
	fmt.Println(then.Hour())
	fmt.Println(then.Minute())
	fmt.Println(then.Second())
	fmt.Println(then.Nanosecond())
	fmt.Println(then.Unix())
	fmt.Println(then.Location())
	fmt.Println(time.Now().Location())

	//Comparison
	fmt.Printf("Is %v before %v?\t: %t\n", time.Now(), then, time.Now().Before(then))
	fmt.Printf("Is %v after %v?\t: %t\n", time.Now(), then, time.Now().After(then))
	fmt.Printf("Is %v same as %v?\t: %t\n", time.Now(), then, time.Now().Equal(then))

	//Subtract times
	difference := time.Now().Sub(then)
	fmt.Println(difference)

	//Add Times
	fmt.Println(then.Add(difference))
	fmt.Println(then.Add(-difference))

	//Time from Unix epoch Jan-01-1970 00:00:00 UTC not including leap seconds
	fmt.Println(time.Unix(0, 0).UTC()) //Seconds, Nanoseconds from epoch

}
