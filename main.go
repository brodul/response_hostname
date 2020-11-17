package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func responseHostname(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	io.WriteString(w, "hostname:"+name)
	fmt.Println("Got request:" + r.RemoteAddr + " " + r.RequestURI)
}

func main() {
	http.HandleFunc("/", responseHostname)
	fmt.Println("Listening on port 8000")
	http.ListenAndServe(":8000", nil)
}
