package common

// IList ...
type IList interface {
	Add(id string, data interface{}) error
	Remove(ids ...string) interface{}
	Size() int
	IsEmpty() bool
	Dump() string
}
