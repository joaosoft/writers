package gowriter

import (
	"io"
	"os"
	"sync"
	"time"

	"fmt"

	"github.com/joaosoft/go-manager/service"
	"github.com/satori/go.uuid"
)

// fileConfig ...
type stdoutConfig struct {
	flushTime time.Duration
}

// StdoutWriter ...
type StdoutWriter struct {
	writer        io.Writer
	config        *stdoutConfig
	queue         gomanager.IList
	formatHandler FormatHandler
	mux           *sync.Mutex
	outOnEmpty    bool
	quit          chan bool
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
				fileWriter.outOnEmpty = true
				if fileWriter.queue.IsEmpty() {
					return
				}

			case <-time.After(fileWriter.config.flushTime):
				for fileWriter.queue.Size() > 0 {
					value := fileWriter.queue.Remove()
					switch value.(type) {
					case []byte:
						stdoutWriter.writer.Write(value.([]byte))
					case Message:
						message := value.(Message)
						if bytes, err := stdoutWriter.formatHandler(message.Prefixes, message.Tags, message.Message, message.Fields); err != nil {
							continue
						} else {
							stdoutWriter.writer.Write(bytes)
						}
					}
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
func (stdoutWriter *StdoutWriter) Write(message []byte) (n int, err error) {
	stdoutWriter.queue.Add(uuid.NewV4().String(), fmt.Sprintf("%s\n", message))
	return 0, nil
}

// SWrite ...
func (stdoutWriter *StdoutWriter) SWrite(prefixes map[string]interface{}, tags map[string]interface{}, message interface{}, fields map[string]interface{}) (n int, err error) {
	stdoutWriter.queue.Add(uuid.NewV4().String(), Message{Prefixes: prefixes, Tags: tags, Message: fmt.Sprintf("%s\n", message), Fields: fields})
	return 0, nil
}
