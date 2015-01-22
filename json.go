//JSON

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Response struct {
	StatusCode    int
	Body          string
	Error         string
	internalCount int
}

type keyedResponse struct {
	StatusCode    int    `status_code`
	Body          string `body`
	Error         string `error`
	internalCount int
}

type jsonHolder struct {
	Num  float64  `json:"num"`
	Strs []string `json:"strs"`
}

// type jsonHolder2 struct {
// 	Page   float64  `json:"num"`
// 	Fruits []string `json:"strs"`
// }

func EncodeJSONExamples() {
	//JSON Encode boolean
	bolB, err := json.Marshal(true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bolB))

	//JSON Encode integer
	intB, err := json.Marshal(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(intB))

	//JSON Encode float
	fltB, err := json.Marshal(3.4500)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fltB))

	//JSON Encode string
	strB, err := json.Marshal("gopher")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(strB))

	//JSON Encode slice ---> JSON array
	slcB, err := json.Marshal([]string{"boo", "bah", "goo", "gah"})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(slcB))

	//JSON Encode map (string => string) ----> JSON object
	mapD := map[string]string{"New York": "Albany", "New South Wales": "Sydney", "Texas": "Austin", "California": "Sacramento", "Victoria": "Melbourne"}
	mapB, err := json.Marshal(mapD)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(mapB))

	//JSON Encode map (string => []string) ----> JSON object
	mapD2 := map[string][]string{
		"Australia": []string{"Sydney", "Melbourne", "Adelaide"},
		"USA":       []string{"New York", "San Francisco", "Los Angeles", "Boston", "Raleigh"},
	}
	mapB2, err := json.Marshal(mapD2)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(mapB2))

	//Encode custom types. Exported fields only, keys are field names by default.
	myResponse := &Response{
		StatusCode:    200,
		Body:          "my body text",
		Error:         "",
		internalCount: 55,
	}
	myResponseB, err := json.Marshal(myResponse)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(myResponseB))

	//Encode custom types. Key names can be tags instead.
	myKeyedResponse := &keyedResponse{
		StatusCode:    200,
		Body:          "my body text",
		Error:         "",
		internalCount: 55,
	}
	myKeyedResponseB, err := json.Marshal(myKeyedResponse)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(myKeyedResponseB))

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}

func DecodeJSONExamples() {

	//Example JSON Data
	jsonData := []byte(`{"num":6.13,"strs":["a","b"]}`)

	//Decode into map
	var data map[string]interface{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		panic(err)
	}
	fmt.Println(data)
	fmt.Println(data["num"].(float64))
	fmt.Println(data["strs"].([]interface{})[0])

	//Decode into custom type
	holder := new(jsonHolder)
	if err := json.Unmarshal(jsonData, &holder); err != nil {
		panic(err)
	}
	fmt.Println(holder)

}

func main() {

	EncodeJSONExamples()

	DecodeJSONExamples()

}
