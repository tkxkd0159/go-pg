package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

import (
	"github.com/tkxkd0159/go-pg/go-basic/nomadlib"
)

var myLogger *log.Logger

func main() {
	fpLog, err := os.OpenFile("./log/info.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer func(fpLog *os.File) {
		err := fpLog.Close()
		if err != nil {

		}
	}(fpLog)

	multiWriter := io.MultiWriter(os.Stdout, fpLog)
	myLogger = log.New(multiWriter, "INFO: ", log.LstdFlags|log.Lshortfile)
	myLogger.Println("Custom logger test")
	fmt.Println("Start")
	nomadlib.WelcomeBasic()
}
