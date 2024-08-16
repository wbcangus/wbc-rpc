package codec

import (
	"bytes"
	"encoding/json"
)

type JsonCodec struct {
}

func (c *JsonCodec) Encode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (c *JsonCodec) Decode(data []byte, v interface{}) error {
	decode := json.NewDecoder(bytes.NewBuffer(data))
	decode.UseNumber()
	return decode.Decode(v)
}
