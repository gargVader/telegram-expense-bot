package log

import (
	"io"
	"log"
	"os"
)

var (
	Logger *log.Logger
)

func Init() {
	file, err := os.OpenFile("data/logs.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	Logger = log.New(io.MultiWriter(file, os.Stdout), "", log.Ldate|log.Ltime)
}

func x() {
	log.Fatal()
}
