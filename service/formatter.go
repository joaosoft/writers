package gowriter

import (
	"encoding/json"
	"fmt"
)

type FormatHandler func(prefixes map[string]interface{}, tags map[string]interface{}, message interface{}, fields map[string]interface{}) ([]byte, error)

func JsonFormatHandler(prefixes map[string]interface{}, tags map[string]interface{}, message interface{}, fields map[string]interface{}) ([]byte, error) {
	if bytes, err := json.Marshal(Message{Prefixes: prefixes, Tags: tags, Message: fmt.Sprint(message), Fields: fields}); err != nil {
		return nil, err
	} else {
		return bytes, nil
	}
}

func TextFormatHandler(prefixes map[string]interface{}, tags map[string]interface{}, message interface{}, fields map[string]interface{}) ([]byte, error) {
	type MessageText struct {
		prefixes interface{}
		tags     interface{}
		message  interface{}
		fields   interface{}
	}

	return []byte(fmt.Sprintf("%+v", MessageText{prefixes: prefixes, tags: tags, message: fmt.Sprint(message), fields: fields})), nil
}
