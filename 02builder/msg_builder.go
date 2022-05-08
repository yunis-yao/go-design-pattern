package builder

import "sync"

//Message 复杂对象的创建
type Message struct {
	Header *Header
	Body   *Body
}
type Header struct {
	SrcAddr  string
	SrcPort  uint64
	DestAddr string
	DestPort uint64
	Items    map[string]string
}
type Body struct {
	Items []string
}

// Message对象的Builder对象
type msgBuilder struct {
	once *sync.Once
	msg  *Message
}

// 返回Builder对象
func Builder() *msgBuilder {
	return &msgBuilder{
		once: &sync.Once{},
		msg:  &Message{Header: &Header{}, Body: &Body{}},
	}
}

// 以下是对Message成员对构建方法
func (b *msgBuilder) WithSrcAddr(srcAddr string) *msgBuilder {
	b.msg.Header.SrcAddr = srcAddr
	return b
}
func (b *msgBuilder) WithSrcPort(srcPort uint64) *msgBuilder {
	b.msg.Header.SrcPort = srcPort
	return b
}
func (b *msgBuilder) WithDestAddr(destAddr string) *msgBuilder {
	b.msg.Header.DestAddr = destAddr
	return b
}
func (b *msgBuilder) WithDestPort(destPort uint64) *msgBuilder {
	b.msg.Header.DestPort = destPort
	return b
}
func (b *msgBuilder) WithHeaderItem(key, value string) *msgBuilder {
	// 保证map只初始化一次
	b.once.Do(func() {
		b.msg.Header.Items = make(map[string]string)
	})
	b.msg.Header.Items[key] = value
	return b
}
func (b *msgBuilder) WithBodyItem(record string) *msgBuilder {
	b.msg.Body.Items = append(b.msg.Body.Items, record)
	return b
}

// 创建Message对象，在最后一步调用
func (b *msgBuilder) Build() *Message {
	return b.msg
}
