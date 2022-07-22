package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Root path")
		io.WriteString(rw, "Hello World!")
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Ops!", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(rw, "%s", data)
	})

	http.HandleFunc("/endpoint", func(rw http.ResponseWriter, _ *http.Request) {
		io.WriteString(rw, "Some endpoint")
	})

	http.ListenAndServe(":9090", nil)
}
