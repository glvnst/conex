package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// command-line option init & defaults
var listenAddr = `0.0.0.0:8000`

// constants
const helpText = `usage: %s

HTTP endpoint reports the user's IP address back to the user

`

// HandleMyIP is an http endpoint that returns the IP address of the requester
func HandleMyIP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// Accepting HTTP GETs only
	if r.Method != http.MethodGet {
		log.Printf("invalid request method! method=%s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Invalid request method\n")
		return
	}

	ip := r.Header.Get(`X-Forwarded-For`)
	if ip == "" {
		ip = r.RemoteAddr
	} else {
		commaIdx := strings.Index(ip, ",")
		if commaIdx > -1 {
			ip = ip[0:commaIdx]
		}
	}

	log.Printf("%s, %s", r.RemoteAddr, r.Header.Get(`X-Forwarded-For`))
	w.Header().Add(`Content-Type`, `application/json`)
	fmt.Fprintf(w, "{\"ip\": \"%s\"}\n", ip)
}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, helpText, os.Args[0])
		flag.PrintDefaults()
	}
	flag.StringVar(&listenAddr, "listen", listenAddr,
		"local address and port to listen on")
	flag.Parse()
}

func main() {
	// setup webhook listener
	http.HandleFunc("/", HandleMyIP)

	// Serve forever
	log.Printf("Starting listener on %s.\n", listenAddr)
	err := http.ListenAndServe(listenAddr, nil)
	log.Fatal(err)
}
