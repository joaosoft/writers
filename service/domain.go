package gowriter

type Message struct {
	Prefixes map[string]interface{} `json:"prefixes,omitempty"`
	Tags     map[string]interface{} `json:"tags,omitempty"`
	Message  interface{}            `json:"message,omitempty"`
	Fields   map[string]interface{} `json:"fields,omitempty"`
}
