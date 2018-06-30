# writer
[![Build Status](https://travis-ci.org/joaosoft/writer.svg?branch=master)](https://travis-ci.org/joaosoft/writer) | [![codecov](https://codecov.io/gh/joaosoft/writer/branch/master/graph/badge.svg)](https://codecov.io/gh/joaosoft/writer) | [![Go Report Card](https://goreportcard.com/badge/github.com/joaosoft/writer)](https://goreportcard.com/report/github.com/joaosoft/writer) | [![GoDoc](https://godoc.org/github.com/joaosoft/writer?status.svg)](https://godoc.org/github.com/joaosoft/writer)

A starting project with writer interface implementations.

###### If i miss something or you have something interesting, please be part of this project. Let me know! My contact is at the end.

## With support for
* file (with queue processing)[1] 
* stdout (with queue processing)[1] [[here]](https://github.com/joaosoft/writer/tree/master/example)

[1] this writer allows you to continue the processing and dispatch the logging  

## Dependecy Management 
>### Dep

Project dependencies are managed using Dep. Read more about [Dep](https://github.com/golang/dep).
* Install dependencies: `dep ensure`
* Update dependencies: `dep ensure -update`


>### Go
```
go get github.com/joaosoft/writer
```

## Interface 
```go
type Writer interface {
	Write(p []byte) (n int, err error)
}
```

## Usage 
This examples are available in the project at [writer/example](https://github.com/joaosoft/writer/tree/master/example)

```go
quit := make(chan bool)
//
// file writer
w := writer.NewFileWriter(
    writer.WithDirectory("./testing"),
    writer.WithFileName("dummy_"),
    writer.WithFileMaxMegaByteSize(1),
    writer.WithFlushTime(time.Second),
    writer.WithQuitChannel(quit),
)

// logger
log := logger.NewLog(
    logger.WithLevel(logger.InfoLevel),
    logger.WithFormatHandler(logger.JsonFormatHandler),
    logger.WithWriter(w)).WithPrefixes(map[string]interface{}{
    "level":   logger.LEVEL,
    "time":    logger.TIME,
    "service": "writer"})

fmt.Printf("send...")
for i := 1; i < 100000; i++ {
    log.Info(fmt.Sprintf("hello number %d\n", i))
}
fmt.Printf("sent!")

// wait one minute to process...
<-time.After(time.Minute * 1)
quit <- true
```

## Known issues


## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com
