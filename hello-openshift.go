package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	helloEnv := os.Getenv("HELLO")

	if helloEnv == "" {
		helloEnv = "hello"
	}

	arg := "openshift: Bom dia INEP, dia 5 começa as inscrições do ENEM!"
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}

	log.Printf("%s %s\n", helloEnv, arg)

	fmt.Fprintf(w, "%s %s\n", helloEnv, arg)
}

func main() {

	http.HandleFunc("/", hello)
	log.Println("Starting server on 0.0.0.0:8080")
	http.ListenAndServe(":8080", nil)
}
