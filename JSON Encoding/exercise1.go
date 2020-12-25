package main

import (
	"encoding/json"
	"log"
	"net/http"
)

/*
	Create a server. Use Encode to send a slice of valies of type transport to the client as JSON.
	Use curl to make a request to the server

*/

// Transport -> Name, Mode, Accomodation
type Transport struct {
	Name         string
	Mode         string
	Accomodation int
}

func main() {
	http.HandleFunc("/encode", TransportEncode)
	err := http.ListenAndServe(":1080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// TransportEncode -> Encoding slice values of type Transport
func TransportEncode(writer http.ResponseWriter, reader *http.Request) {
	ship := Transport{Name: "Titanic", Mode: "Water", Accomodation: 9500000}
	bus := Transport{Name: "Bus", Mode: "Land", Accomodation: 100}
	cycle := Transport{Name: "Bicycle", Mode: "Land", Accomodation: 2}

	transports := []Transport{ship, bus, cycle}

	err := json.NewEncoder(writer).Encode(transports)
	if err != nil {
		log.Println("Encoding Unsuccessfull\n", err)
	}
}
