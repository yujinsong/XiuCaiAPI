package utils

import (
	"github.com/XiuCai/XiuCaiAPI/utils/go-logging"
	"log"
	"os"
)

var (
	Logger  *logging.Logger = nil
	logPath string          = ""
)

func init() {

	//curdir, err := filepath.Rel("/mnt/logs", "/log")
	//if err != nil {
	//	log.Fatal("get log save path failed", err)
	//	os.Exit(1)
	//}
	logPath = "/var/log/"
	//logPath = "./"
}

func CreateLogger(fileName string) {
	logger := logging.MustGetLogger("logger")
	//logging.SetLevel(logging.WARNING, "logger")
	//
	format := logging.MustStringFormatter(
		`%{color}%{time:2006-01-02 15:04:05.000} %{shortfile} %{longfunc} >>> %{level:.4s} %{id:04d} %{message}%{color:reset}`,
	)

	logFilePath := logPath + fileName
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		log.Fatal("open log file error", err)
	}


	backend1 := logging.NewLogBackend(logFile, "", 0)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)

	// Only errors and more severe messages should be sent to backend1
	backend1Formatter := logging.NewBackendFormatter(backend1, format)
	backend2Formatter := logging.NewBackendFormatter(backend2, format)
//	backend3Formatter := logging.NewBackendFormatter(backend1Leveled, format)

	//logging.SetBackend(backend1Leveled, backend3Formatter)
	logging.SetBackend(backend1Formatter, backend2Formatter)
	//logging.SetLevel(logging.INFO, "logger")

	Logger = logger

}
