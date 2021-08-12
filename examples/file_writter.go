package main

import (
	"fmt"
	"time"
	"github.com/joaosoft/writers"
)

func ExampleFileWriter() {
	quit := make(chan bool)
	fmt.Println(":: FILE WRITER")

	//
	// file
	fileWriter := writers.NewFileWriter(
		writers.WithFileDirectory("./testing"),
		writers.WithFileName("dummy_"),
		writers.WithFileMaxMegaByteSize(1),
		writers.WithFileFlushTime(time.Second*5),
		writers.WithFileQuitChannel(quit),
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
