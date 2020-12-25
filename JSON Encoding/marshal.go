package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Id int
	FirstName string
	LastName string
}

func main() {

	person1 := Person{Id: 1, FirstName: "James", LastName: "Monroe"}
	person2 := Person{Id: 2, FirstName: "Jane", LastName: "Doe"}
	person3 := Person{Id: 3, FirstName: "Mary", LastName: "Jane"}

	employees := []Person{person1, person2, person3}

	bytes, err := json.Marshal(employees)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(string(bytes))
}
