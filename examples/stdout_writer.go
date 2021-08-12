package main

import (
	"fmt"
	"github.com/joaosoft/writers"
	"time"
)

func ExampleStdoutWriter() {
	quit := make(chan bool)
	fmt.Println(":: STDOUT WRITER")

	//
	// json
	stdoutWriter := writers.NewStdoutWriter(
		writers.WithStdoutFlushTime(time.Second*5),
		writers.WithStdoutQuitChannel(quit),
		writers.WithStdoutFormatHandler(writers.JsonFormatHandler),
	)

	fmt.Println("send...")
	dummy := make(map[string]interface{})
	for i := 1; i < 100000; i++ {
		stdoutWriter.SWrite(dummy, dummy, fmt.Sprintf("hello number %d", i), dummy, dummy)
	}
	fmt.Println("sent!")

	// wait one minute to process...
	<-time.After(time.Second * 10)

	fmt.Println("quitting...")
	quit <- true
}
