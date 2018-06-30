package main

import (
	"fmt"
	"time"
	"writer"
)

func ExampleFileWriter() {
	quit := make(chan bool)
	fmt.Println(":: FILE WRITER")

	//
	// file
	fileWriter := writer.NewFileWriter(
		writer.WithFileDirectory("./testing"),
		writer.WithFileName("dummy_"),
		writer.WithFileMaxMegaByteSize(1),
		writer.WithFileFlushTime(time.Second*5),
		writer.WithFileQuitChannel(quit),
	)

	fmt.Println("send...")
	for i := 1; i < 100000; i++ {
		fileWriter.Write([]byte(fmt.Sprintf("hello number %d", i)))
	}
	fmt.Println("sent!")

	// wait one minute to process...
	<-time.After(time.Second * 10)

	fmt.Println("quitting...")
	quit <- true

}
