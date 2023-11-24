package main

import (
	"fmt"
	"io"
	"net/http"
)

func server1func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is first runserve.")
}

func server2func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is second runserve")
	io.WriteString(w, r.URL.Host)
}
