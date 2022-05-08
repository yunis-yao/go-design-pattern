package builder

import (
	"testing"
)

func Test_Nomarl_Create_MsgObj(t *testing.T) {
	// 一般情况下创建对象
	message := Message{
		Header: &Header{
			SrcAddr:  "192.168.0.1",
			SrcPort:  1234,
			DestAddr: "192.168.0.2",
			DestPort: 8080,
			Items:    make(map[string]string),
		},
		Body: &Body{
			Items: make([]string, 0),
		},
	}

	// 需要知道对象的实现细节
	message.Header.Items["contents"] = "application/json"
	message.Body.Items = append(message.Body.Items, "record1")
	message.Body.Items = append(message.Body.Items, "record2")

	t.Log(message)
}

func TestMessageBuilder(t *testing.T) {
	// 使用消息建造者进行对象创建
	message := Builder().
		WithSrcAddr("192.168.0.1").
		WithSrcPort(1234).
		WithDestAddr("192.168.0.2").
		WithDestPort(8080).
		WithHeaderItem("contents", "application/json").
		WithBodyItem("record1").
		WithBodyItem("record2").
		Build()
	if message.Header.SrcAddr != "192.168.0.1" {
		t.Errorf("expect src address 192.168.0.1, but actual %s.", message.Header.SrcAddr)
	}

	if message.Body.Items[0] != "record1" {
		t.Errorf("expect body item0 record1, but actual %s.", message.Body.Items[0])
	}
}
