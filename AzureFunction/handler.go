package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Kapparina/gogoprojecttemplates/AzureFunction/Hello"
)

var (
	greetingHandler = Hello.GreetingHandler{}
)

func main() {
	listenAddr := ":7071"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		if val == "" {
			val = "7071"
		}
		listenAddr = ":" + val
	}
	http.Handle("/api/Hello/", greetingHandler)
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
