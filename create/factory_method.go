package create

type DpFactory interface {
	NewDp(dpId uint16, dpVal string) (Dper, error)
}

// A 客户实现了 NilDp 实现类
type NilDpFactory func(dpId uint16, dpVal string)

func (n NilDpFactory) NewDp(dpId uint16, _ string) (Dper, error) {
	return NewNilDp(dpId)
}

// B 客户实现了 FileDp 实现类
type FileDpFactory func(dpId uint16, dpVal string)

func (n FileDpFactory) NewDp(dpId uint16, dpVal string) (Dper, error) {
	return NewFileDp(dpId, dpVal)
}
