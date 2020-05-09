package myevent

// 事件
type Event struct {
	Name       string
	Parameters map[string]interface{}
}

// 事件监听者
type Listener interface {
	Handle(event *Event)
}

// 带有 ch 的 listener
type ListenerWithCh struct {
	listener Listener
	ch       chan *Event
}

// 创建一个带 chan 的监听者
func newListenerWithCh(listener Listener, chLen int32) *ListenerWithCh {
	return &ListenerWithCh{
		listener: listener,
		ch:       make(chan *Event, chLen),
	}
}

// 事件派发器
type Dispatcher struct {
	// 注册的事件监听者
	listeners map[string]*[]*ListenerWithCh
	// 事件处理器的 chan 的缓冲区长度
	chLen int32
}

// 创建一个事件派发器
func NewDisPatcher(chLen int32) *Dispatcher {
	listeners := make(map[string]*[]*ListenerWithCh)
	return &Dispatcher{
		listeners: listeners,
		chLen:     chLen,
	}
}

// 添加监听者
func (d *Dispatcher) AddListener(eventName string, listener Listener) {
	listenerWithChs, ok := d.listeners[eventName]
	if !ok {
		newListenerWithChs := make([]*ListenerWithCh, 0)
		d.listeners[eventName] = &newListenerWithChs
		listenerWithChs = &newListenerWithChs
	}

	// 插入新的 listeners
	listenerWithCh := newListenerWithCh(listener, d.chLen)
	*listenerWithChs = append(*listenerWithChs, listenerWithCh)
	// 启动新 goroutine 等待从 chan 里获取 event
	go d.handler(listenerWithCh)
}

func (d *Dispatcher) handler(listenerWithChan *ListenerWithCh) {
	// 等待事件
	for e := range listenerWithChan.ch {
		// 启动新的 goroutine 来处理事件
		go listenerWithChan.listener.Handle(e)
	}
}

// 派发一个任务
func (d *Dispatcher) DispatchEvent(event *Event) {
	if event == nil {
		return
	}
	listenerWithChans, ok := d.listeners[event.Name]
	if !ok {
		return
	}
	for _, listenerWithChan := range *listenerWithChans {
		listenerWithChan.ch <- event
	}
}
