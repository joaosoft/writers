package main

import (
	"encoding/json"
	"fmt"
	"time"
	"github.com/joaosoft/writers"
)

func ExampleStdoutWriterWithStruct() {
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
	dummy2 := struct {
		Name string `json:"name"`
	}{
		Name: "joao",
	}

	body, _ := json.Marshal(dummy2)
	fmt.Printf(string(body))

	stdoutWriter.Write(body)
	fmt.Println("sent!")

	fmt.Println("quitting...")
	quit <- true
}
