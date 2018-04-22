package main

import (
	"fmt"
	"go-writer/service"
	"time"
)

func RunWriteAsJsonToFile() {
	quit := make(chan bool)
	fmt.Println(":: FILE WRITER")

	//
	// file
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

	fmt.Println("quitting...")
	quit <- true

}
