package main

import (
	"encoding/json"
	"fmt"
	"go-writer/service"
	"time"
)

func RunWriteOthers() {
	quit := make(chan bool)
	fmt.Println(":: STDOUT WRITER")

	//
	// json
	stdoutWriter := gowriter.NewStdoutWriter(
		gowriter.WithStdoutFlushTime(time.Second*5),
		gowriter.WithStdoutQuitChannel(quit),
		gowriter.WithStdoutFormatHandler(gowriter.JsonFormatHandler),
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
