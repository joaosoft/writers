package gowriter

import (
	"encoding/json"
	"fmt"
	"time"

	logger "github.com/joaosoft/go-log/service"
)

type FormatHandler func(level logger.Level, message logger.Message) ([]byte, error)

func JsonFormatHandler(level logger.Level, message logger.Message) ([]byte, error) {
	addSystemInfo(level, &message)
	if bytes, err := json.Marshal(message); err != nil {
		return nil, err
	} else {
		return bytes, nil
	}
}

func TextFormatHandler(level logger.Level, message logger.Message) ([]byte, error) {
	type MessageText struct {
		prefixes interface{}
		tags     interface{}
		message  interface{}
		fields   interface{}
	}

	addSystemInfo(level, &message)
	return []byte(fmt.Sprintf("%+v", MessageText{prefixes: message.Prefixes, tags: message.Tags, message: message.Message, fields: message.Fields})), nil
}

func addSystemInfo(level logger.Level, message *logger.Message) {
	// special prefixes keys
	prefixes := make(map[string]interface{}, len(message.Prefixes))
	for key, value := range message.Prefixes {
		switch value {
		case logger.LEVEL:
			value = level.String()
		case logger.TIME:
			value = time.Now().Format("2006-01-02 15:04:05:06")
		}
		prefixes[key] = value
	}
	message.Prefixes = prefixes
}
