package main

import (
	"fmt"
	"go-writer/service"
	"time"
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

	fmt.Println("send...")
	for i := 1; i < 100000; i++ {
		fileWriter.Write([]byte(fmt.Sprintf("hello number %d", i)))
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
		gowriter.WithStdoutFormatHandler(gowriter.JsonFormatHandler),
	)

	fmt.Println("send...")
	dummy := make(map[string]interface{})
	for i := 1; i < 100000; i++ {
		stdoutWriter.SWrite(dummy, dummy, fmt.Sprintf("hello number %d", i), dummy)
	}
	fmt.Println("sent!")

	// wait one minute to process...
	<-time.After(time.Second * 10)

	fmt.Println("QUITTING...")
	quit <- true
}
