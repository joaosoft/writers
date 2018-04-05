package gowriter

import (
	"encoding/json"
	"fmt"
)

type FormatHandler func(message Message) ([]byte, error)

func JsonFormatHandler(message Message) ([]byte, error) {
	if bytes, err := json.Marshal(message); err != nil {
		return nil, err
	} else {
		return bytes, nil
	}
}

func TextFormatHandler(message Message) ([]byte, error) {
	type MessageText struct {
		prefixes interface{}
		tags     interface{}
		message  interface{}
		fields   interface{}
	}

	return []byte(fmt.Sprintf("%+v", MessageText{prefixes: message.Prefixes, tags: message.Tags, message: message.Message, fields: message.Fields})), nil
}
