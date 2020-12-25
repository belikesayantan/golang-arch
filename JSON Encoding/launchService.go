package main

import "net/http"

func main() {
	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)

	http.ListenAndServe(":8080", nil)
}

func foo(writer http.ResponseWriter, reader *http.Request) {

}

func bar(writer http.ResponseWriter, reader *http.Request) {
	
}
