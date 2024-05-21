package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/dalibormesaric/go-key-value-api/internal/server"
)

var port = flag.Int("p", 9000, "port number")

func main() {
	flag.Parse()
	fmt.Printf("Running on port %d\n", *port)

	s := server.NewHTTPServer(*port)
	log.Fatal(s.ListenAndServe())
}
