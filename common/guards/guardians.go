package guards

import (
	"errors"
	"mutating-webhook/common/log"
)

var logger = log.Log


func FailOnError(err error, msg string){
	if err != nil {
		fail(err)
	}
}



func NotImplemented(){
	fail(errors.New("function not implemented"))
}



func FailOnLessOrEqualToZero(num int, msg string) {
	if num <= 0 {
		fail(errors.New(msg))
	}
}


func FailOnEmptyString(str string, msg string ){
	if str == "" {
		fail(errors.New(msg))
	}
}


func FailOnNil(entity interface{}, msg string){
	if entity == nil {
		fail(errors.New(msg))
	}
}


func fail(err error){
	logger.Panic().Msgf("%+v", err.Error())
}