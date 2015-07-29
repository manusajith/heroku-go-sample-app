package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// Parse port from command-line parameters
	port := flag.String("port", "8080", "HTTP Port to listen on")
	flag.Parse()

	// Start our Server
	log.Println("Starting Server on", *port)
	http.HandleFunc("/", all)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func all(w http.ResponseWriter, r *http.Request) {
	content := "Hello World..!!\n Email: neo@codingarena.in\n Github: http://github.com/manusajith"
	w.Write([]byte(content))
}
