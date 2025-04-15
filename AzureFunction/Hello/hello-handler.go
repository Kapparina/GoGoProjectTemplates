package Hello

import (
	"fmt"
	"net/http"
)

type GreetingHandler struct{}

func (g GreetingHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if err := greet(writer, request); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func greet(w http.ResponseWriter, r *http.Request) error {
	var message string
	if name := r.URL.Query().Get("name"); len(name) == 0 {
		message = "Hello, World!"
	} else {
		message = fmt.Sprintf("Hello, %s!", name)
	}
	_, err := w.Write([]byte(message))
	return err
}
