package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Human struct {
	Id int
	FirstName string
	LastName string
}

func main() {

	person1 := Human{Id: 1, FirstName: "James", LastName: "Monroe"}
	person2 := Human{Id: 2, FirstName: "Jane", LastName: "Doe"}
	person3 := Human{Id: 3, FirstName: "Mary", LastName: "Jane"}

	employees := []Human{person1, person2, person3}

	/* Converting Data into Json */
	byteSlice, err := json.Marshal(employees)
	if err != nil {
		log.Panic(err)
	}

	//fmt.Println(string(byteSlice))

	/* Converting Json to data */
	var NewEmployees []Human

	err = json.Unmarshal(byteSlice, &NewEmployees)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(NewEmployees)
}

