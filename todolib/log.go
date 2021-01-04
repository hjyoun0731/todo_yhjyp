package todolib

import (
	"log"
	"os"
)

// MyLogger for logging
var MyLogger *log.Logger

//FpLog logfile
var FpLog *os.File
var err error

func init() {
	FpLog, err = os.OpenFile("log/logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	MyLogger = log.New(FpLog, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// MakeLog logging func
func MakeLog(msg string) {
	MyLogger.Print(msg)
	log.Println(msg)
}
