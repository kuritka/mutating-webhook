package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"time"

	"mutating-webhook/common/guards"
	"mutating-webhook/common/log"
)

var logger = log.Log
const port = 8443

func handleJson(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello %q", html.EscapeString(r.URL.Path))
}

func handleMutate(w http.ResponseWriter, r *http.Request) {

	// read the body / request
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		guards.FailOnError(err, "reading Body")
	}

	fmt.Println(string(body))

	// mutate the request
	// body, err = mutant.Mutate(body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		guards.FailOnError(err, "mutate")
	}

	// and write it back
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}


func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/json", handleJson)
	mux.HandleFunc("/mutate", handleMutate)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%v", port),
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1048576
	}

	logger.Info().Msgf("Listening on %v:",port)
	err := s.ListenAndServeTLS("/etc/webhook/certs/cert.pem","/etc/webhook/certs/key.pem" )
	guards.FailOnError(err,"server stopped")
}