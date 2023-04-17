package main

import (
	"fmt"
	"log"
	"os"
)

var LOGFILE = "./deferlog.log"

func log1(l *log.Logger) {
	l.Println("--log1------")
	defer l.Println("--log1-----")

	for i := 0; i < 10; i++ {
		l.Println(i)
	}
}

func log2(l *log.Logger) {
	l.Println("log2------")
	defer l.Println("------log2")

	for i := 0; i < 10; i++ {
		l.Println(i)
	}
}

func main() {
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	logDefer := log.New(f, "logDefer", log.LstdFlags)
	logDefer.Println("Start logging.....")

	log1(logDefer)
	log2(logDefer)
}
