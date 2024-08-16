package codec

// Codec 序列化接口
type Codec interface {
	// Encode 序列化
	Encode(i interface{}) ([]byte, error)
	// Decode 反序列化
	Decode(data []byte, i interface{}) error
}
