package create

import "testing"

func TestRecvService_doWork(t *testing.T) {
	// 初始化接收数据的业务对象
	//service := NewRecvService(new(AModelDeviceProtocol))
	service := NewRecvService(new(BModelDeviceProtocol))
	service.DoWork("", "")
}
