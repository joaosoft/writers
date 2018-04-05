package gowriter

import (
	"io"
	"os"
	"sync"
	"time"

	"github.com/joaosoft/go-manager/service"
	"github.com/satori/go.uuid"
)

// fileConfig ...
type stdoutConfig struct {
	flushTime time.Duration
}

// StdoutWriter ...
type StdoutWriter struct {
	writer     io.Writer
	config     *stdoutConfig
	queue      gomanager.IList
	mux        *sync.Mutex
	outOnEmpty bool
	quit       chan bool
}

// NewStdoutWriter ...
func NewStdoutWriter(options ...StdoutWriterOption) *StdoutWriter {
	stdoutWriter := &StdoutWriter{
		queue:  gomanager.NewQueue(gomanager.WithMode(gomanager.FIFO)),
		writer: os.Stdout,
		mux:    &sync.Mutex{},
		config: &stdoutConfig{},
		quit:   make(chan bool),
	}
	stdoutWriter.Reconfigure(options...)
	stdoutWriter.process()

	return stdoutWriter
}

func (stdoutWriter *StdoutWriter) process() error {
	go func(fileWriter *StdoutWriter) {
		for {
			select {
			case <-fileWriter.quit:
				if fileWriter.queue.IsEmpty() {
					return
				} else {
					fileWriter.outOnEmpty = true
				}

			case <-time.After(fileWriter.config.flushTime):
				for fileWriter.queue.Size() > 0 {
					stdoutWriter.writer.Write(fileWriter.queue.Remove().([]byte))
				}

				if fileWriter.queue.IsEmpty() && stdoutWriter.outOnEmpty {
					return
				}
			}
		}
	}(stdoutWriter)
	return nil
}

// Write ...
func (stdoutWriter StdoutWriter) Write(message []byte) (n int, err error) {
	stdoutWriter.queue.Add(uuid.NewV4().String(), message)
	return 0, nil
}
