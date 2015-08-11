package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// BuildHandler run build
func BuildHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Fprint(w, "start build")
		// TODO: run cache.go and update build_status
		// TODO: run build and update build_status
	}
}

// BuildStatusHandler response build_status
func BuildStatusHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Is build_status JSON? or Text?
	status, err := ioutil.ReadFile("./build_status.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, string(status))
}

func main() {
	http.HandleFunc("/build", BuildHandler)
	http.HandleFunc("/build/status", BuildStatusHandler)
	http.ListenAndServe(":3000", nil)
}
