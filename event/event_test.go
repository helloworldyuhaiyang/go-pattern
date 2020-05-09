package event

import (
	"fmt"
	"testing"
	"time"
)

type DpChangeRule struct {
}

func (t *DpChangeRule) Change(event *Event) {
	dpId := event.Params["dpId"]
	oldVal := event.Params["oldVal"]
	newVal := event.Params["newVal"]

	fmt.Println(dpId, oldVal, newVal)
}

func TestDispatcher_DispatchEvent(t *testing.T) {
	rule := DpChangeRule{}
	var eventCallback EventCallback = rule.Change
	//获取分派器单例
	dispatcher := SharedDispatcher()
	dispatcher.AddEventListener("dpChange", &eventCallback)

	//随便弄个事件携带的参数，我把参数定义为一个map
	params := make(map[string]interface{})
	params["dpId"] = 1000
	params["oldVal"] = true
	params["newVal"] = false
	//创建一个事件对象
	event := CreateEvent("dpChange", params)

	dispatcher.DispatchEvent(event)

	time.Sleep(time.Second * 2)
}
