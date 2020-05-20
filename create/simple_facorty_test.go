package create

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewBooleanDp(t *testing.T) {
	type args struct {
		id     uint16
		rawVal interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *BooleanDp
		wantErr bool
	}{
		{
			name: "testStringTrue",
			args: args{1, "true"},
			want: &BooleanDp{baseDp{
				Id:       1,
				Type:     "boolean",
				RawValue: true,
				ValueStr: "true",
			}},
			wantErr: false,
		},
		{
			name: "testStringFalse",
			args: args{2, "false"},
			want: &BooleanDp{baseDp{
				Id:       2,
				Type:     "boolean",
				RawValue: false,
				ValueStr: "false",
			}},
			wantErr: false,
		},
		{
			name: "testBoolTrue",
			args: args{3, true},
			want: &BooleanDp{baseDp{
				Id:       3,
				Type:     "boolean",
				RawValue: true,
				ValueStr: "true",
			}},
			wantErr: false,
		},
		{
			name: "testBoolFalse",
			args: args{3, false},
			want: &BooleanDp{baseDp{
				Id:       3,
				Type:     "boolean",
				RawValue: false,
				ValueStr: "false",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBooleanDp(tt.args.id, tt.args.rawVal)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBooleanDp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBooleanDp() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewByteArrayDp(t *testing.T) {
	type args struct {
		id     uint16
		rawVal interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *ByteArrayDp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewByteArrayDp(tt.args.id, tt.args.rawVal)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewByteArrayDp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewByteArrayDp() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewIntDp(t *testing.T) {
	type args struct {
		id      uint16
		typeStr string
		rawVal  interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *IntDp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewIntDp(tt.args.id, tt.args.typeStr, tt.args.rawVal)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewIntDp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIntDp() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNilDp(t *testing.T) {
	type args struct {
		id uint16
	}
	tests := []struct {
		name string
		args args
		want *NilDp
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := NewNilDp(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNilDp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewStringDp(t *testing.T) {
	type args struct {
		id     uint16
		rawVal interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *StringDp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewStringDp(tt.args.id, tt.args.rawVal)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewStringDp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStringDp() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_test(t *testing.T) {
	a := 1.1
	if isInteger(a) {
		fmt.Printf("a should not integer,a=%f\n", a)
	}
	b := 100.0
	if !isInteger(b) {
		fmt.Printf("a should  integer,b=%f\n", b)
	}
}

func TestIntDp_CompareTo(t *testing.T) {
	dpInt80, _ := NewInt8Dp(300, 1)
	dpInt81, _ := NewInt8Dp(300, 1)
	dpInt82, _ := NewInt8Dp(300, 2)

	dpBoolean0, _ := NewBooleanDp(300, true)
	dpBoolean1, _ := NewBooleanDp(300, true)
	dpBoolean2, _ := NewBooleanDp(300, false)

	dpString0, _ := NewStringDp(300, "hello")
	dpString1, _ := NewStringDp(300, "hello")
	dpString2, _ := NewStringDp(300, "nihao")

	dpByteArray0, _ := NewByteArrayDp(300, "01020304")
	dpByteArray1, _ := NewByteArrayDp(300, "01020304")
	dpByteArray2, _ := NewByteArrayDp(300, "0102030405")

	type fields struct {
		Dper
	}
	type args struct {
		other Dper
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "int8_equal",
			fields: fields{
				dpInt80,
			},
			args:    args{dpInt81},
			want:    0,
			wantErr: false,
		},
		{
			name: "int8_not_equal",
			fields: fields{
				dpInt80,
			},
			args:    args{dpInt82},
			want:    -1,
			wantErr: false,
		},
		{
			name: "boolean_equal",
			fields: fields{
				dpBoolean0,
			},
			args:    args{dpBoolean1},
			want:    0,
			wantErr: false,
		},
		{
			name: "boolean_not_equal",
			fields: fields{
				dpBoolean0,
			},
			args:    args{dpBoolean2},
			want:    1,
			wantErr: false,
		}, {
			name: "string_equal",
			fields: fields{
				dpString0,
			},
			args:    args{dpString1},
			want:    0,
			wantErr: false,
		},
		{
			name: "string_not_equal",
			fields: fields{
				dpString0,
			},
			args:    args{dpString2},
			want:    -1,
			wantErr: false,
		},
		{
			name: "byteArray_equal",
			fields: fields{
				dpByteArray0,
			},
			args:    args{dpByteArray1},
			want:    0,
			wantErr: false,
		},
		{
			name: "byteArray_not_equal1",
			fields: fields{
				dpByteArray0,
			},
			args:    args{dpByteArray2},
			want:    -5,
			wantErr: false,
		},
		{
			name: "byteArray_not_equal2",
			fields: fields{
				dpByteArray2,
			},
			args:    args{dpByteArray0},
			want:    5,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.Dper.CompareTo(tt.args.other)
			if (err != nil) != tt.wantErr {
				t.Errorf("CompareTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CompareTo() got = %v, want %v", got, tt.want)
			}
		})
	}
}
