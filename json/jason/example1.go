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

	// string to byte array
	exampleJSONBytes := []byte(exampleJSON)
	j, err1 := jason.NewObjectFromBytes(exampleJSONBytes)

	log.Printf("%v-> %s (err: %s)\n", exampleJSON, exampleJSONBytes, err1)
	name, errname := j.GetString("name")
	age, errage := j.GetNumber("age")
	log.Println("name:", name, ", err:", errname)
	log.Println("age:", age, ", err:", errage)

	children, errchildren := j.GetStringArray("children")

	for i, child := range children {
		log.Println(i, "th, child:", child, "err", errchildren)
	}

}
