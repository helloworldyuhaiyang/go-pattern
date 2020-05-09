package myevent

import (
	"fmt"
	"github.com/iothink/dp-service/protocol/dp"
	"testing"
	"time"
)

type DpChangeListener struct {
}

func NewDpChangeListener() *DpChangeListener {
	return &DpChangeListener{}
}

func (d *DpChangeListener) Handle(event *Event) {
	fmt.Println("recv event", event)
}

func TestDispatcher_DispatchEvent(t *testing.T) {
	disPatcher := NewDisPatcher(10)
	disPatcher.AddListener("dpChange", NewDpChangeListener())

	parameters := make(map[string]interface{})
	parameters["deviceId"] = "dev1234"
	parameters["dp"], _ = dp.NewInt8Dp(1, 123)
	event := &Event{
		Name:       "dpChange",
		Parameters: parameters,
	}
	disPatcher.DispatchEvent(event)

	time.Sleep(time.Second)
}
