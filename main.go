package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"time"

	"mutating-webhook/common/guards"

	"github.com/rs/zerolog"
)

var Log *zerolog.Logger

const port = 8443

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello %q", html.EscapeString(r.URL.Path))
}

func handleMutate(w http.ResponseWriter, r *http.Request) {

	Log.Info().Msg("handle mutate")
	// read the body / request
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		Log.Err(err)
	}

	// mutate the request
	fmt.Println(string(body))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err)

	}

	// and write it back
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}


func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/mutate", handleMutate)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%v", port),
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1048576
	}

	Log.Info().Msgf("Listening on %v:",port)
	err := s.ListenAndServeTLS("/etc/webhook/certs/cert.pem","/etc/webhook/certs/key.pem" )
	guards.FailOnError(err,"server stopped")
}