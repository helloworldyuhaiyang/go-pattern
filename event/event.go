package event

import (
	//"fmt"
	"unsafe"
)

type Dispatcher struct {
	listeners map[string]*Chain
}

type Chain struct {
	chs       []chan *Event
	callbacks []*EventCallback
}

func createEventChain() *Chain {
	return &Chain{chs: []chan *Event{}, callbacks: []*EventCallback{}}
}

type Event struct {
	eventName string
	Params    map[string]interface{}
}

func CreateEvent(eventName string, params map[string]interface{}) *Event {
	return &Event{eventName: eventName, Params: params}
}

type EventCallback func(*Event)

var _instance *Dispatcher

func SharedDispatcher() *Dispatcher {
	if _instance == nil {
		_instance = &Dispatcher{}
		_instance.Init()
	}

	return _instance
}

func (d *Dispatcher) Init() {
	d.listeners = make(map[string]*Chain)
}

func (d *Dispatcher) AddEventListener(eventName string, callback *EventCallback) {
	eventChain, ok := d.listeners[eventName]
	if !ok {
		eventChain = createEventChain()
		d.listeners[eventName] = eventChain
	}

	exist := false
	//fmt.Println("add len:", len(eventChain.callbacks))
	for _, item := range eventChain.callbacks {
		a := *(*int)(unsafe.Pointer(item))
		b := *(*int)(unsafe.Pointer(callback))
		//fmt.Println("add", a, b)
		if a == b {
			exist = true
			break
		}
	}

	if exist {
		return
	}

	ch := make(chan *Event)

	eventChain.chs = append(eventChain.chs[:], ch)
	eventChain.callbacks = append(eventChain.callbacks[:], callback)

	go d.handler(ch, callback)
}

func (d *Dispatcher) handler(ch chan *Event, callback *EventCallback) {
	for {
		event := <-ch
		//fmt.Println("event out:", eventName, event, ch)
		if event == nil {
			break
		}
		go (*callback)(event)
	}
}

func (d *Dispatcher) RemoveEventListener(eventName string, callback *EventCallback) {
	eventChain, ok := d.listeners[eventName]
	if !ok {
		return
	}

	var ch chan *Event
	exist := false
	key := 0
	for k, item := range eventChain.callbacks {
		a := *(*int)(unsafe.Pointer(item))
		b := *(*int)(unsafe.Pointer(callback))
		//fmt.Println("remove", a, b)
		if a == b {
			exist = true
			ch = eventChain.chs[k]
			key = k
			break
		}
	}

	if exist {
		//fmt.Printf("remove listener: %s\n", eventName)
		//fmt.Println("chan: ", ch)
		ch <- nil

		eventChain.chs = append(eventChain.chs[:key], eventChain.chs[key+1:]...)
		eventChain.callbacks = append(eventChain.callbacks[:key], eventChain.callbacks[key+1:]...)
		//fmt.Println(len(eventChain.chs))
	}
}

func (d *Dispatcher) DispatchEvent(event *Event) {
	eventChain, ok := d.listeners[event.eventName]
	if ok {
		////fmt.Printf("dispatch event: %s\n", event.eventName)
		for _, chEvent := range eventChain.chs {
			chEvent <- event
		}
	}
}
