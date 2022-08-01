package handlers

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	log.Println("Root path")
	io.WriteString(rw, "Hello!")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Ops!", http.StatusBadRequest)
		return
	}
	if data != nil {
		fmt.Fprintf(rw, " %s", data)
	}
}
