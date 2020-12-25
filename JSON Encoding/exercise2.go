package main

import (
	"encoding/json"
	"log"
	"net/http"
)

/*
	Create a server and decode given below json into appropriate date structures

	[{"Country": "India"}, {"Country": "China"}]
*/

// Country -> Name
type Country struct {
	Name string
}

func main() {
	http.HandleFunc("/decode", CountryDecode)

	err := http.ListenAndServe(":1080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// CountryDecode -> decode json into appropriate data structures
func CountryDecode(w http.ResponseWriter, r *http.Request) {
	var country []Country

	err := json.NewDecoder(r.Body).Decode(&country)
	if err != nil {
		log.Println("Decoding Unsuccessfull:\n", err)
	}

	log.Println("Country: ", country)
}
