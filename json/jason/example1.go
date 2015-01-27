package main

import (
	"github.com/antonholmquist/jason"
	"log"
)

func main() {

	exampleJSON := `{
		"name" : "Walter White",
		"age":51,
		"children":[
			"junior",
			"holly"
		],
		"other":{
			"occupation" : "chemist",
			"years":23
		}
	}`

	exampleJSONBytes := []byte(exampleJSON)

	j, _ := jason.NewObjectFromBytes(exampleJSONBytes)

	name, _ := j.GetString("name")
	age, _ := j.GetNumber("age")
	log.Println("name:", name)
	log.Println("age:", age)
}
