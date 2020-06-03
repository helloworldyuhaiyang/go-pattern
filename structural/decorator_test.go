package structural

import (
	"fmt"
	"testing"
)

func Test_Decorate(t *testing.T) {
	var a StringHandler
	a = func(s string) string {
		fmt.Println(s)
		return s
	}

	// 两次装饰
	decorate := LogDecorate(ReplaceDecorate(a, "hello", "你好"))

	_ = decorate("hello world")
}
