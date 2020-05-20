package create

import (
	"encoding/hex"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// 数据点应该满足的接口
type Dper interface {
	GetId() uint16
	GetType() string
	// 真实类型
	GetRawValue() interface{}
	// 为了写日志写数据库方便 字符串化的类型
	GetValueStr() string
	// 全比较 id,type,val
	CompareTo(other Dper) (int, error)
	// 只比较 type, val
	CompareToVal(other Dper) (int, error)
}

// dp 点类型
//noinspection ALL
type DpType struct {
	TypeStr string
	TypeInt uint8
}

var (
	Nil       = DpType{TypeStr: "nil", TypeInt: 0}
	Boolean   = DpType{TypeStr: "boolean", TypeInt: 1}
	Int8      = DpType{TypeStr: "int8", TypeInt: 2}
	Int16     = DpType{TypeStr: "int16", TypeInt: 3}
	Int32     = DpType{TypeStr: "int32", TypeInt: 4}
	Int64     = DpType{TypeStr: "int64", TypeInt: 5}
	String    = DpType{TypeStr: "string", TypeInt: 6}
	ByteArray = DpType{TypeStr: "byte[]", TypeInt: 7}
	File      = DpType{TypeStr: "file", TypeInt: 8}
)

type baseDp struct {
	Id   uint16
	Type string
	// 存储真实原数据
	RawValue interface{}
	// 这个主要用来存数据库，打日志等
	ValueStr string
}

func (b *baseDp) GetId() uint16 {
	return b.Id
}

func (b *baseDp) GetType() string {
	return b.Type
}

func (b *baseDp) GetRawValue() interface{} {
	return b.RawValue
}

func (b *baseDp) GetValueStr() string {
	return b.ValueStr
}

// type id 不同的话，直接报错，不需要比较
func (b *baseDp) CompareTo(other Dper) (int, error) {
	if b.Type != other.GetType() {
		return 0, errors.New(fmt.Sprintf("b.Type:%v is not eq other.Type:%v", b.Type, other.GetType()))
	}

	if b.Id != other.GetId() {
		return 0, errors.New(fmt.Sprintf("b.Id:%v is not eq other.Id:%v", b.Id, other.GetId()))
	}
	return 0, nil
}

// 创建 dper 的简单工厂
func NewDp(dpId uint16, dpType string, dpVal interface{}) (Dper, error) {
	switch dpType {
	case Nil.TypeStr:
		if nilDp, err := NewNilDp(dpId); err != nil {
			return nil, err
		} else {
			return nilDp, nil
		}
	case Boolean.TypeStr:
		if booleanDp, err := NewBooleanDp(dpId, dpVal); err != nil {
			return nil, err
		} else {
			return booleanDp, nil
		}
	case Int8.TypeStr, Int16.TypeStr, Int32.TypeStr, Int64.TypeStr:
		if intDp, err := NewIntDp(dpId, dpType, dpVal); err != nil {
			return nil, err
		} else {
			return intDp, nil
		}
	case String.TypeStr:
		if stringDp, err := NewStringDp(dpId, dpVal); err != nil {
			return nil, err
		} else {
			return stringDp, nil
		}
	case ByteArray.TypeStr:
		if arrayDp, err := NewByteArrayDp(dpId, dpVal); err != nil {
			return nil, err
		} else {
			return arrayDp, nil
		}
	case File.TypeStr:
		if fileDp, err := NewFileDp(dpId, dpVal); err != nil {
			return nil, err
		} else {
			return fileDp, nil
		}
	default:
		return nil, errors.New("invalid dp type")
	}
}

////////////////////////// NilDp //////////////////////////////
type NilDp struct {
	baseDp
}

func NewNilDp(id uint16) (*NilDp, error) {
	nilDp := new(NilDp)
	nilDp.Id = id
	nilDp.Type = Nil.TypeStr
	nilDp.RawValue = nil
	nilDp.ValueStr = "nil"

	return nilDp, nil
}

func (b *NilDp) CompareTo(other Dper) (int, error) {
	return b.baseDp.CompareTo(other)
}

func (b *NilDp) CompareToVal(_ Dper) (int, error) {
	return 0, nil
}

////////////////////////// BooleanDp //////////////////////////////
type BooleanDp struct {
	baseDp
}

func NewBooleanDp(id uint16, rawVal interface{}) (*BooleanDp, error) {
	booleanDp := new(BooleanDp)
	booleanDp.Id = id
	booleanDp.Type = Boolean.TypeStr
	switch rawVal.(type) {
	case bool:
		if rawVal == true {
			booleanDp.RawValue = true
			booleanDp.ValueStr = "true"
		} else {
			booleanDp.ValueStr = "false"
			booleanDp.RawValue = false
		}
	case string:
		strVal, _ := rawVal.(string)
		if strVal == "true" {
			booleanDp.RawValue = true
			booleanDp.ValueStr = "true"
		} else if strVal == "false" {
			booleanDp.RawValue = false
			booleanDp.ValueStr = "false"
		} else {
			return nil, errors.New("error strVal:" + strVal)
		}
	default:
		return nil, errors.New("value type:" + typeof(rawVal) + " can not convert to bool dp")
	}

	return booleanDp, nil
}

func (b *BooleanDp) CompareTo(other Dper) (int, error) {
	r, err := b.baseDp.CompareTo(other)
	if err != nil || r != 0 {
		return r, err
	}
	return b.CompareToVal(other)
}
func (b *BooleanDp) CompareToVal(other Dper) (int, error) {
	if b.GetRawValue() == true {
		if other.GetRawValue() == true {
			return 0, nil
		} else {
			return 1, nil
		}
	} else {
		if other.GetRawValue() == true {
			return -1, nil
		} else {
			return 0, nil
		}
	}
}

////////////////////////// IntDp //////////////////////////////

// int 类型的数据点
type IntDp struct {
	baseDp
}

func NewIntDp(id uint16, typeStr string, rawVal interface{}) (*IntDp, error) {
	intDp := new(IntDp)
	intDp.Id = id
	intDp.Type = typeStr
	// 把值统一转为 int
	i, err := convertInt64(rawVal)
	if err != nil {
		return nil, err
	}
	// 值的字符串表示
	intDp.ValueStr = strconv.Itoa(int(i))
	// 设置 RawValue
	switch typeStr {
	case Int8.TypeStr:
		intDp.RawValue = int8(i)
	case Int16.TypeStr:
		intDp.RawValue = int16(i)
	case Int32.TypeStr:
		intDp.RawValue = int32(i)
	case Int64.TypeStr:
		intDp.RawValue = i
	default:
		return nil, errors.New("invalid int typeStr")
	}
	return intDp, nil
}

func convertInt64(val interface{}) (int64, error) {
	switch val.(type) {
	case int:
		return int64(val.(int)), nil
	case uint:
		return int64(val.(uint)), nil
	case int8:
		return int64(val.(int8)), nil
	case uint8:
		return int64(val.(uint8)), nil
	case int16:
		return int64(val.(int16)), nil
	case uint16:
		return int64(val.(uint16)), nil
	case int32:
		return int64(val.(int32)), nil
	case uint32:
		return int64(val.(uint32)), nil
	case int64:
		return val.(int64), nil
	case uint64:
		// 不考虑越界问题
		return int64(val.(uint64)), nil
	case string:
		str, _ := val.(string)
		if intVal, e := strconv.Atoi(str); e != nil {
			return 0, e
		} else {
			return int64(intVal), nil
		}
	case float64:
		float, _ := val.(float64)
		if isInteger(float) {
			return int64(float), nil
		} else {
			return 0, errors.New("float64 is not integer")
		}
	case float32:
		float, _ := val.(float32)
		if isInteger(float64(float)) {
			return int64(float), nil
		} else {
			return 0, errors.New("float64 is not integer")
		}
	default:
		return 0, errors.New("value type:" + typeof(val) + " can not convert to int dp")
	}
}

func NewInt8Dp(id uint16, val int8) (*IntDp, error) {
	return NewIntDp(id, Int8.TypeStr, val)
}

func NewInt16Dp(id uint16, val int16) (*IntDp, error) {
	return NewIntDp(id, Int16.TypeStr, val)
}

func NewInt32Dp(id uint16, val int32) (*IntDp, error) {
	return NewIntDp(id, Int32.TypeStr, val)
}

func NewInt64Dp(id uint16, val int64) (*IntDp, error) {
	return NewIntDp(id, Int64.TypeStr, val)
}

func (b *IntDp) CompareTo(other Dper) (int, error) {
	r, err := b.baseDp.CompareTo(other)
	if err != nil || r != 0 {
		return r, err
	}
	return b.CompareToVal(other)
}

func (b *IntDp) CompareToVal(other Dper) (int, error) {
	// 都转为 int 然后减操作
	// 统一转为 int64 计算
	x, err := convertInt64(b.GetRawValue())
	if err != nil {
		return 0, err
	}
	y, err := convertInt64(other.GetRawValue())
	if err != nil {
		return 0, err
	}
	return int(x - y), nil
}

////////////////////////// StringDp //////////////////////////////

// string 类型数据点
type StringDp struct {
	baseDp
}

func NewStringDp(id uint16, rawVal interface{}) (*StringDp, error) {
	strDp := new(StringDp)
	strDp.Id = id
	strDp.Type = String.TypeStr
	strDp.RawValue = rawVal
	switch rawVal.(type) {
	case string:
		str, _ := rawVal.(string)
		strDp.ValueStr = str
	default:
		return nil, errors.New("value type:" + typeof(rawVal) + " can not convert to string dp")
	}

	return strDp, nil
}

func (b *StringDp) CompareTo(other Dper) (int, error) {
	r, err := b.baseDp.CompareTo(other)
	if err != nil || r != 0 {
		return r, err
	}
	return b.CompareToVal(other)
}

func (b *StringDp) CompareToVal(other Dper) (int, error) {
	// 字符串对比
	return strings.Compare(b.GetValueStr(), other.GetValueStr()), nil
}

////////////////////////// ByteArrayDp //////////////////////////////

// []byte 类型的数据点
type ByteArrayDp struct {
	baseDp
}

func NewByteArrayDp(id uint16, rawVal interface{}) (*ByteArrayDp, error) {
	byteArrayDp := new(ByteArrayDp)
	byteArrayDp.Id = id
	byteArrayDp.Type = ByteArray.TypeStr
	switch rawVal.(type) {
	case []byte:
		if byteArray, ok := rawVal.([]byte); !ok {
			byteArrayDp.RawValue = rawVal
			byteArrayDp.ValueStr = hex.EncodeToString(byteArray)
		}
	case string:
		hexString, _ := rawVal.(string)
		if decodeString, err := hex.DecodeString(hexString); err != nil {
			return nil, err
		} else {
			byteArrayDp.RawValue = decodeString
			byteArrayDp.ValueStr = hexString
		}
	default:
		return nil, errors.New("value type:" + typeof(rawVal) + " can not convert []byte dp")
	}

	return byteArrayDp, nil
}

func (b *ByteArrayDp) CompareTo(other Dper) (int, error) {
	r, err := b.baseDp.CompareTo(other)
	if err != nil || r != 0 {
		return r, err
	}
	return b.CompareToVal(other)
}

func (b *ByteArrayDp) CompareToVal(other Dper) (int, error) {
	// 转数组
	xArray := b.GetRawValue().([]byte)
	yArray := other.GetRawValue().([]byte)
	xLen := len(xArray)
	yLen := len(yArray)
	var minLen = 0
	if xLen < yLen {
		minLen = xLen
	} else {
		minLen = yLen
	}

	// 一个一个相减对比
	var i = 0
	for i = 0; i < minLen; i++ {
		xVal := xArray[i]
		yVal := yArray[i]
		if xVal-yVal == 0 {
			continue
		} else {
			return int(xVal - yVal), nil
		}
	}
	// 长的 大
	if xLen > yLen {
		return int(xArray[i] - 0), nil
	} else if yLen > xLen {
		return -int(yArray[i]), nil
	} else {
		return 0, nil
	}
}

//////////////////////////// 文件类型的 dp 点 //////////////////////////
type FileDp struct {
	baseDp
}

func NewFileDp(id uint16, rawVal interface{}) (*FileDp, error) {
	fileDp := new(FileDp)
	fileDp.Id = id
	fileDp.Type = File.TypeStr
	fileDp.RawValue = rawVal
	switch rawVal.(type) {
	case string:
		str, _ := rawVal.(string)
		fileDp.ValueStr = str
	default:
		return nil, errors.New("value type:" + typeof(rawVal) + " can not convert to file dp")
	}

	return fileDp, nil
}

func (b *FileDp) CompareTo(other Dper) (int, error) {
	r, err := b.baseDp.CompareTo(other)
	if err != nil || r != 0 {
		return r, err
	}
	// 字符串对比
	return b.CompareToVal(other)
}

func (b *FileDp) CompareToVal(other Dper) (int, error) {
	// 字符串对比
	return strings.Compare(b.GetValueStr(), other.GetValueStr()), nil
}

// 获取类型名
func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func isInteger(a float64) bool {
	if a-float64(int64(a)) == 0 {
		return true
	} else {
		return false
	}
}
