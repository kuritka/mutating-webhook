package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mutating-webhook/common/flags"
	"mutating-webhook/webhooks/mutlabel"
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
	envLabels="LABEL_MAP"
)

var labels map[string]string

func handleMutate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		guards.HttpThrowServerError(w,err, "can't read request body")
		return
	}
	defer r.Body.Close()


	// verify the content type is accurate
	if contentType := r.Header.Get("Content-Type"); contentType != "application/json" {
		guards.HttpThrowError(w,http.StatusInternalServerError,"Content-Type=%s, expect application/json", contentType)
		return
	}

	body, err = mutlabel.Mutate(body, labels)
	if err != nil {
		guards.HttpThrowServerError(w,err,"can't mutate request")
		return
	}

	// and write it back
	if _, err = w.Write(body); err != nil {
		guards.HttpThrowServerError(w, err, "can't send AdmissionReview")
	}
}


func readLabelConfig() map[string]string {
	labels := make(map[string]string)
	labelBase64 := flags.MustGetStringFlagFromEnv(envLabels)
	labelsJson, err := base64.StdEncoding.DecodeString(labelBase64)
	guards.FailOnError(err,"invalid labels %v", labelBase64)
	err = json.Unmarshal(labelsJson, &labels)
	guards.FailOnError(err,"invalid labels %v", string(labelsJson))
	return labels
}

func main() {
	labels = readLabelConfig()

	mux := http.NewServeMux()
	mux.HandleFunc("/mutate-labels", handleMutate)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%v", port),
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1048576
	}

	logger.Info().Msgf("Listening on %v:",port)
	logger.Info().Msgf("%v", labels)
	err := s.ListenAndServeTLS(certPath,keyPath )
	guards.FailOnError(err,"server stopped")
}