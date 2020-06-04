package structural

import "fmt"

//  手机输入电压,手机针对这个接口编程
type MobileInput interface {
	Input5V() int
}

// 家里交流电
type FamilyPower struct {
	power int
}

func NewFamilyPower() *FamilyPower {
	return &FamilyPower{power: 220}
}

func (f *FamilyPower) OutPut() int {
	return f.power
}

// 适配器,实现接口 MobileInput
type Adapter struct {
	FamilyPower
}

func NewAdapter(familyPower FamilyPower) *Adapter {
	return &Adapter{FamilyPower: familyPower}
}

// 输出 5V
func (a *Adapter) Input5V() int {
	fmt.Printf("我是适配器, 输入:%d,输出:%d", a.power, 5)
	return 5
}
