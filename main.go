package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)


func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello %q", html.EscapeString(r.URL.Path))
}

func handleMutate(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w,"handle mutate")
	// read the body / request
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err)
	}

	// mutate the request
	fmt.Println(string(body))

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err)
	}

	// and write it back
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

type params struct {
	cert string
	key string
}


func main() {

	//p := new(params)
	//flag.StringVar(&p.cert, "tlsCertFile", "/etc/webhook/certs/cert.pem", "File containing the x509 Certificate for HTTPS.")
	//flag.StringVar(&p.key, "tlsKeyFile", "/etc/webhook/certs/key.pem", "File containing the x509 private key to --tlsCertFile.")
	//flag.Parse()

	fmt.Println("Started")
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/mutate", handleMutate)

	s := &http.Server{
		Addr:           ":8443",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1048576
	}

	log.Println("Listening on :8443")
	err := s.ListenAndServeTLS("/etc/webhook/certs/cert.pem","/etc/webhook/certs/key.pem" )
	//err := s.ListenAndServeTLS("", "")
	if err != nil {
		log.Panic(err)
	}
}