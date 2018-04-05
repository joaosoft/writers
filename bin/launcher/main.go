package main

import (
	"fmt"
	"go-writer/service"
	"time"

	logger "github.com/joaosoft/go-log/service"
)

func main() {
	quit := make(chan bool)
	fmt.Println(":: FILE WRITER")

	//
	// file fileWriter
	fileWriter := gowriter.NewFileWriter(
		gowriter.WithFileDirectory("./testing"),
		gowriter.WithFileName("dummy_"),
		gowriter.WithFileMaxMegaByteSize(1),
		gowriter.WithFileFlushTime(time.Second*5),
		gowriter.WithFileQuitChannel(quit),
	)

	// logger
	log := logger.NewLog(
		logger.WithLevel(logger.InfoLevel),
		logger.WithFormatHandler(logger.JsonFormatHandler),
		logger.WithWriter(fileWriter)).WithPrefixes(map[string]interface{}{
		"level":   logger.LEVEL,
		"time":    logger.TIME,
		"service": "go-Writer"})

	fmt.Println("send...")
	for i := 1; i < 100000; i++ {
		log.Info(fmt.Sprintf("hello number %d\n", i))
	}
	fmt.Println("sent!")

	// wait one minute to process...
	<-time.After(time.Second * 10)

	fmt.Println("QUITTING...")
	quit <- true

	fmt.Println(":: STDOUT WRITER")

	//
	// stdout fileWriter
	stdoutWriter := gowriter.NewStdoutWriter(
		gowriter.WithStdoutFlushTime(time.Second*5),
		gowriter.WithStdoutQuitChannel(quit),
	)

	// logger
	log = logger.NewLog(
		logger.WithLevel(logger.InfoLevel),
		logger.WithFormatHandler(logger.JsonFormatHandler),
		logger.WithWriter(stdoutWriter)).WithPrefixes(map[string]interface{}{
		"level":   logger.LEVEL,
		"time":    logger.TIME,
		"service": "go-Writer"})

	fmt.Println("send...")
	for i := 1; i < 100000; i++ {
		log.Info(fmt.Sprintf("hello number %d\n", i))
	}
	fmt.Println("sent!")

	// wait one minute to process...
	<-time.After(time.Second * 10)

	fmt.Println("QUITTING...")
	quit <- true
}
