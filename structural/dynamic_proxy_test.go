package structural

import (
	"bou.ke/monkey"
	"fmt"
	"os"
	"testing"
)

func Test_DynamicProxy(t *testing.T) {
	monkey.Patch(fmt.Println, func(a ...interface{}) (n int, err error) {
		return fmt.Fprintln(os.Stdout, "你好")
	})
	fmt.Println("what the hell?") // what the *bleep*?
}
