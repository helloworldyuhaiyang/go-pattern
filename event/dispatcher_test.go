package event

import (
	"fmt"
	"testing"
	"time"
)

// 关注 dpChange 类型的观察者
type DpChangeListener struct {
}

func NewDpChangeListener() *DpChangeListener {
	return &DpChangeListener{}
}

func (d *DpChangeListener) Handle(event *Event) {
	fmt.Println("recv event", event)
}

// 单测代码
func TestDispatcher_DispatchEvent(t *testing.T) {
	disPatcher := NewDisPatcher(10)
	disPatcher.AddListener("dpChange", NewDpChangeListener())

	parameters := make(map[string]interface{})
	parameters["deviceId"] = "dev1234"
	parameters["dpVal"] = "123"
	event := &Event{
		Name:       "dpChange",
		Parameters: parameters,
	}
	disPatcher.DispatchEvent(event)

	time.Sleep(time.Second)
}
