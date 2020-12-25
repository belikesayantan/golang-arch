package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Animal struct {
	ID int
	Type string
	Name string
	Color string
}


func main() {
	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func foo(writer http.ResponseWriter, reader *http.Request) {
	cow := Animal{ID: 1, Type: "Mammal", Name: "Bos Taurus", Color: "White"}
	//dog := Animal{ID: 2, Type: "Mammal", Name: "Canis lupus familiaris", Color: "Yellow"}

	err := json.NewEncoder(writer).Encode(cow)
	if err != nil {
		// Here, log.Panic(err) will display the error as well as shutdown the running server.
		// So, to prevent that, just log the error
		log.Println("Encoding unsuccessful", err)
	}
}

func bar(writer http.ResponseWriter, reader *http.Request) {
	//cat := Animal{ID: 3, Type: "Mammal", Name: "Felis catus", Color: "Brown"}
}
