package gowriter

import (
	"bufio"
	"bytes"
	"os"
	"sync"
	"time"

	"fmt"

	"encoding/binary"

	"github.com/joaosoft/go-manager/service"
	uuid "github.com/satori/go.uuid"
)

// fileConfig ...
type fileConfig struct {
	directory   string
	fileName    string
	fileMaxSize int64
	flushTime   time.Duration
}

// StdoutWriter ...
type FileWriter struct {
	writer     *bufio.Writer
	config     *fileConfig
	queue      gomanager.IList
	mux        *sync.Mutex
	outOnEmpty bool
	quit       chan bool
}

// NewStdoutWriter ...
func NewFileWriter(options ...FileWriterOption) *FileWriter {
	fileWriter := &FileWriter{
		queue:  gomanager.NewQueue(gomanager.WithMode(gomanager.FIFO)),
		mux:    &sync.Mutex{},
		config: &fileConfig{},
		quit:   make(chan bool),
	}
	fileWriter.Reconfigure(options...)
	fileWriter.process()

	return fileWriter
}

func (fileWriter *FileWriter) process() error {
	if _, err := os.Stat(fileWriter.config.directory); os.IsNotExist(err) {
		if err = os.Mkdir(fileWriter.config.directory, 0777); err != nil {
			return err
		}
	}

	go func(fileWriter *FileWriter) {
		var tmpLogFileName string
		var logMessage []byte
		for {
			select {
			case <-fileWriter.quit:
				if fileWriter.queue.IsEmpty() {
					return
				} else {
					fileWriter.outOnEmpty = true
				}

			case <-time.After(fileWriter.config.flushTime):
				fileWriter.mux.Lock()
				defer fileWriter.mux.Unlock()

			newFile:
				tmpLogFileName = fmt.Sprintf("%s/%s%s", fileWriter.config.directory, fileWriter.config.fileName, time.Now().Format("2006.01.02 15.04.05.06"))
				file, err := os.OpenFile(tmpLogFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
				checkError(err, fmt.Sprintf("error opening file %s: %s", tmpLogFileName, err), file)

				fileSize, _ := file.Stat()
				maxSize := fileWriter.config.fileMaxSize - fileSize.Size()
				buffer := bytes.NewBuffer(make([]byte, 0))

				for fileWriter.queue.Size() > 0 {
					logMessage = fileWriter.queue.Remove().([]byte)

					if int64(binary.Size(buffer.Bytes())+binary.Size(logMessage)) <= maxSize {
						buffer.Write(logMessage)
					} else {
						if _, err := file.Write(buffer.Bytes()); err != nil {
							checkError(err, fmt.Sprintf("error writing file %s: %s", tmpLogFileName, err), file)
						}
						file.Close()
						goto newFile
					}
				}

				if _, err := file.Write(buffer.Bytes()); err != nil {
					checkError(err, fmt.Sprintf("error flushing to file %s: %s", tmpLogFileName, err), file)
				}
				file.Close()

				if fileWriter.queue.IsEmpty() && fileWriter.outOnEmpty {
					return
				}
			}
		}
	}(fileWriter)
	return nil
}

// Write ...
func (fileWriter FileWriter) Write(message []byte) (n int, err error) {
	fileWriter.queue.Add(uuid.NewV4().String(), message)
	return 0, nil
}
