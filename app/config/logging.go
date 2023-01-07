package config

import (
	"io"
	"log"
	"os"
)

func Logging(source string) {
	logfile, err := os.OpenFile(source, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	mutiLogfile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(mutiLogfile)
}
