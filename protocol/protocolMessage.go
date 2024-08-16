package protocol

type ProtocolMessage struct {
	Header
	body interface{}
}

// Header 消息头占 17 个字节
type Header struct {
	Magic      byte   // 魔法值，保证安全性
	Version    byte   // 版本号
	Type       byte   // 消息类型：请求或者响应
	Serializer byte   // 序列化方式
	Status     byte   // 响应的状态
	RequestId  uint64 // 请求 id
	BodyLength uint32 // 消息体长度
}
