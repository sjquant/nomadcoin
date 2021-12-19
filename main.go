package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sjquant/nomadcoin/rest"
	"github.com/sjquant/nomadcoin/web"
)

func usage() {
	fmt.Printf("Welcome to 노마드 코인\n\n")
	fmt.Printf("Please use the following commands:\n\n")
	fmt.Printf("web:        Start the HTML Explorer\n")
	fmt.Printf("rest: 	    Start the REST API (recommended)\n\n")
	os.Exit(0)
}

func main() {

	if len(os.Args) < 2 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'web' and 'rest'")

	flag.Parse()

	switch *mode {
	case "web":
		web.Start(*port)
		break
	case "rest":
		rest.Start(*port)
		break
	default:
		usage()
	}
}
