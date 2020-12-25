package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Animal Type
type Animal struct {
	ID    int
	Type  string
	Name  string
	Color string
}

func main() {
	http.HandleFunc("/encode", encode)
	http.HandleFunc("/decode", decode)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

/* curl localhost:8080/encode */
func encode(writer http.ResponseWriter, reader *http.Request) {
	cow := Animal{ID: 1, Type: "Mammal", Name: "Bos Taurus", Color: "White"}
	//dog := Animal{ID: 2, Type: "Mammal", Name: "Canis lupus familiaris", Color: "Yellow"}
	//cat := Animal{ID: 3, Type: "Mammal", Name: "Felis catus", Color: "Brown"}

	err := json.NewEncoder(writer).Encode(cow)
	if err != nil {
		// Here, log.Panic(err) will display the error as well as shutdown the running server.
		// So, to prevent that, just log the error
		log.Println("Encoding unsuccessful", err)
	}
}

/* curl -XGET -H "Content-type: application/json" -d '{"ID":1,"Type":"Mammal","Name":"Bos Taurus","Color":"White"}' 'localhost:8080/decode' */
func decode(writer http.ResponseWriter, reader *http.Request) {
	var animal Animal
	err := json.NewDecoder(reader.Body).Decode(&animal)
	if err != nil {
		log.Println("Decoding unsuccessful", err)
	}

	log.Println("Animal: ", animal)

	/*
		To use the decode facility use curl command, use the online curl command line builder
	*/

}
