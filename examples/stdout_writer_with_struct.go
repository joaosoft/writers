package main

import (
	"encoding/json"
	"fmt"
	"time"
	"writer"
)

func ExampleStdoutWriterWithStruct() {
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
