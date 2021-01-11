package main

import (
	"encoding/base64"
	"fmt"
)

/*
	HTTP Basic Authentication
	-------------------------

	*	Basic Authentication part of the http
		-- send username/password with every request

		-- uses authorization header & keyword "Basic"
			*	put "username:password" together

			*	convert them to base64
				- put generic binary data into printable form
				- base64 is reversible
					* only use with https, not http

					*	use basic authentication to login
*/

func main() {
	var user, pass string
	_, _ = fmt.Scanln(&user)
	_, _ = fmt.Scanln(&pass)

	fmt.Println("Basic " + base64.StdEncoding.EncodeToString([]byte(user+":"+pass)))
}
