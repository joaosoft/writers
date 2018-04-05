package main

import (
	"fmt"
	"go-log/service"
	"go-writer/service"
	"time"
)

func main() {
	//
	// stdout fileWriter
	quit := make(chan bool)
	stdoutWriter := gowriter.NewStdoutWriter(
		gowriter.WithStdoutFlushTime(time.Second*10),
		gowriter.WithStdoutQuitChannel(quit),
	)

	//
	// log to json
	fmt.Println(":: LOG JSON")
	log := golog.NewLog(
		golog.WithLevel(golog.InfoLevel),
		golog.WithFormatHandler(gowriter.JsonFormatHandler),
		golog.WithWriter(stdoutWriter)).
		With(
			map[string]interface{}{"level": golog.LEVEL, "time": golog.TIME},
			map[string]interface{}{"service": "log"},
			map[string]interface{}{"name": "joão"})

	// logging...
	start := time.Now()
	sum := 0
	for i := 0; i < 100000; i++ {
		log.Infof("MESSAGE %d", i+1)
		sum += 1
	}
	elapsed := time.Since(start)
	log.Infof("ELAPSED TIME: %s", elapsed.String())

	<-time.After(time.Second * 30)
	quit <- true
}
