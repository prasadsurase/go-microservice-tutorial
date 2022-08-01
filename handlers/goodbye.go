package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// io.WriteString(rw, "Good bye")
	rw.Write([]byte("Good bye"))
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Ops!", http.StatusBadRequest)
		return
	}
	if data != nil {
		fmt.Fprintf(rw, " %s", data)
	}
}
