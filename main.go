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
const
(
	port = 8443
	//I don't expect any changes in volumeMounts that's why I'm keeping paths as constants
	certPath ="/etc/webhook/certs/cert.pem"
	keyPath= "/etc/webhook/certs/key.pem"
)

func handleMutate(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	guards.HttpFailOnError(w,err,"can't read request body")


	// verify the content type is accurate
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		guards.HttpThrowError(w,http.StatusUnsupportedMediaType,"Content-Type=%s, expect application/json", contentType)
	}

	body, err = mutate.Mutate(body)
	guards.HttpFailOnError(w,err,"can't mutate request body")


	// and write it back
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(body)
	guards.HttpFailOnError(w,err,"can't send AdmissionReview")
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
	err := s.ListenAndServeTLS(certPath,keyPath )
	guards.FailOnError(err,"server stopped")
}