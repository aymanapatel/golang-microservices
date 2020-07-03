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

	d, err := ioutil.ReadAll(r.Body)
	
	h.l.Println("Println: Hello ", string(d))  // Use h.l.Println() (instead of log.Println) for better testability and future DI work. log.Println("Hello world")

	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
	}

	fmt.Fprintf(rw, "ResponseWriter:  Hello %s \n", d)
}
