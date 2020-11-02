package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello Tamilarasan")
		d, _ := ioutil.ReadAll(r.Body)

		//	log.Printf("Data %s\n", d) ---> Request
		fmt.Fprintf(w, "Hello  %s", d) // ---->Respons
	})

	http.HandleFunc("/bye", func(http.ResponseWriter, *http.Request) {

		log.Println("goodbye Tamilarasan")

	})

	http.ListenAndServe(":9000", nil)

}
