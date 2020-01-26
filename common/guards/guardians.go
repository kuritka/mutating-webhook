package guards

import (
	"mutating-webhook/common/log"
	"net/http"
)

var logger = log.Log


func HttpFailOnError(w http.ResponseWriter, err error, message string, v ...interface{}){
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Panic().Err(err).Msgf(message, v)
	}
}

func HttpThrowError(w http.ResponseWriter, httpCode int, message string, v ...interface{}){
	http.Error(w, message, httpCode)
	logger.Panic().Msgf(message, v)
}


func FailOnError( err error, message string, v ...interface{}){
	if err != nil {
		logger.Panic().Err(err).Msgf(message, v)
	}
}
