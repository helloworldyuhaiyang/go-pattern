package create

import "fmt"

// 我有协议定义如下
// 传输层 用来传输数据的。 例如 mqtt
// 解释层 用来解析数据成应用数据. 例如 json
// 应用层  应用层业务实现需要的数据类型 我们这里就是 Dper

// 传输层
type Transporter interface {
	Send([]byte) (int, error)
	Recv() ([]byte, error)
}

// 解释层
type PreCoder interface {
	EnCode(dper ...Dper) []byte
	DeCode([]byte) []Dper
}

/////////////// 我们有设备类型为 A 的产品使用 mqtt,json,dp 来通信 ////////////////

// mqtt 作为传输层
type MqttTransporter struct {
}

func (m *MqttTransporter) Send(bytes []byte) (int, error) {
	fmt.Println("MqttTransporter send data")
	return 0, nil
}

func (m *MqttTransporter) Recv() ([]byte, error) {
	// 从 mqtt 链接获取数据
	fmt.Println("MqttTransporter Recv data")
	return make([]byte, 0), nil
}

// json 作为解释层编解码
type JsonPreCoder struct {
}

func (j *JsonPreCoder) EnCode(dper ...Dper) []byte {
	fmt.Println("JsonPreCoder Encode")
	return make([]byte, 0)
}

func (j *JsonPreCoder) DeCode(bytes []byte) []Dper {
	fmt.Println("JsonPreCoder DeCode")
	return make([]Dper, 0)
}

// A 型号的产品的通信方式
type AModelDeviceProtocol struct {
}

func (p *AModelDeviceProtocol) GetTransporter() Transporter {
	return new(MqttTransporter)
}

func (p *AModelDeviceProtocol) GetPreCoder() PreCoder {
	return new(JsonPreCoder)
}

//////////// 有设备类型 B 型号的产品使用 udp,tlv 的方式来通信 ///////////
// udp 作为传输层
type UdpTransporter struct {
}

func (u *UdpTransporter) Send(bytes []byte) (int, error) {
	fmt.Println("UdpTransporter send data")
	return 0, nil
}

func (u *UdpTransporter) Recv() ([]byte, error) {
	// 从 udp 链接获取数据
	fmt.Println("UdpTransporter Recv data")
	return make([]byte, 0), nil
}

// tlv 作为解释层
type TlvPreCoder struct {
}

func (t *TlvPreCoder) EnCode(dper ...Dper) []byte {
	fmt.Println("TlvPreCoder Encode")
	return make([]byte, 0)
}

func (t *TlvPreCoder) DeCode(bytes []byte) []Dper {
	fmt.Println("TlvPreCoder DeCode")
	return make([]Dper, 0)
}

// B 型号的产品的通信方式
type BModelDeviceProtocol struct {
}

func (p *BModelDeviceProtocol) GetTransporter() Transporter {
	return new(UdpTransporter)
}

func (p *BModelDeviceProtocol) GetPreCoder() PreCoder {
	return new(TlvPreCoder)
}

////////////////////////// 抽象工厂 ///////////////////////////////

// 如果我们的业务代码调用的代码针对 AModelDeviceProtocol 实现编程了，那么我们想加入新的 BModelDeviceProtocol 就必须修改代码
// 所以我们要面向抽象编程。不管什么型号设备上报数据,都需 Transporter 和 PreCoder。

// 定义处理不同类似上报数据的抽象工厂
type ProtocolFactory interface {
	GetTransporter() Transporter
	GetPreCoder() PreCoder
}

// AModelDeviceProtocol BModelDeviceProtocol 都实现了ProtocolFactory抽象接口 我们针对这个接口编程。
// 只要调用注入 不同的实现类就好了。
type RecvService struct {
	p ProtocolFactory
}

// 带参数的对象的构造需要实现 New 方法
func NewRecvService(p ProtocolFactory) *RecvService {
	return &RecvService{p: p}
}

func (s *RecvService) DoWork(a, b string) {
	recv, _ := s.p.GetTransporter().Recv()
	dpers := s.p.GetPreCoder().DeCode(recv)
	// do somthing
	fmt.Println(dpers)
}
