package main

import (
	"fmt"
	"io/ioutil"
	"mutating-webhook/mutate"
	"net/http"
	"time"

	"mutating-webhook/common/guards"
	"mutating-webhook/common/log"
)

var logger = log.Log
const port = 8443


func handleMutate(w http.ResponseWriter, r *http.Request) {

	// read the body / request
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Err(err).Msg( "can't reading request body")
	}

	fmt.Printf("REQUEST BODY :\n%v\n\n", string(body))
	body, err = mutate.Mutate(body)
	fmt.Printf("REPOSNSE BODY :\n%v\n\n", string(body))


	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Err(err).Msg( "can't mutate request body")
	}

	// and write it back
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Err(err).Msg( "can't send AdmissionReview")
	}
}


func main() {

	mux := http.NewServeMux()
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