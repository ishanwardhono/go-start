package json

import (
	"bytes"
	"encoding/json"
	"io"
)

type Decoder interface {
	Decode(obj interface{}) error
}

func NewDecoder(reader io.Reader) *json.Decoder {
	dec := json.NewDecoder(reader)
	dec.UseNumber()
	return dec
}

func Decode(r io.Reader, v interface{}) error {
	return NewDecoder(r).Decode(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return Decode(bytes.NewReader(data), v)
}

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
