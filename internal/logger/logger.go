package logger

import (
	"log"
	"os"
)

const (
	DEBUG = iota
	INFO
	WARNING
	ERROR
	CRITICAL
	AUDIT
)

var (
	file     *os.File
	logger   *log.Logger
	LogLevel int
)

func NewLogger(logfile string) {
	var err error
	file, err = os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logger = log.New(file, "", log.Ldate|log.Ltime)
}

func Log(level int, message string) {
	if level < LogLevel {
		return
	}

	prefix := ""
	switch level {
	case DEBUG:
		prefix = "[DEBUG] "
	case INFO:
		prefix = "[INFO] "
	case WARNING:
		prefix = "[WARN] "
	case ERROR:
		prefix = "[ERR] "
	case CRITICAL:
		prefix = "[CRIT] "
	case AUDIT:
		prefix = "[AUDIT] "
	default:
		prefix = "[UNK] "
	}

	logger.SetPrefix(prefix)
	logger.Println(message)
}

func Close() {
	file.Close()
}
