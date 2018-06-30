package main

import (
	"fmt"
	"writer"
	"time"
)

func ExampleStdoutWriter() {
	quit := make(chan bool)
	fmt.Println(":: STDOUT WRITER")

	//
	// json
	stdoutWriter := writer.NewStdoutWriter(
		writer.WithStdoutFlushTime(time.Second*5),
		writer.WithStdoutQuitChannel(quit),
		writer.WithStdoutFormatHandler(writer.JsonFormatHandler),
	)

	fmt.Println("send...")
	dummy := make(map[string]interface{})
	for i := 1; i < 100000; i++ {
		stdoutWriter.SWrite(dummy, dummy, fmt.Sprintf("hello number %d", i), dummy)
	}
	fmt.Println("sent!")

	// wait one minute to process...
	<-time.After(time.Second * 10)

	fmt.Println("quitting...")
	quit <- true
}
