package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Hello https://
type Hello struct {
	l *log.Logger
}

//NewHello https://
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	h.l.Println("Decompsoited with go-mods: Hello world. ") // Use h.l.Println() for better testability and future DI work. log.Println("Hello world")
	d, err := ioutil.ReadAll(r.Body)
	//log.Println("Hello birthday to %s ", d) // IP: AYM, OP: [65 95 .. ..]
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
	}

	fmt.Fprintf(rw, "Decomposited Hello with go-mods:  %s \n", d)
}
