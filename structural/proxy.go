package structural

import (
	"fmt"
)

func main() {
	saveUser(&user{"sivan"})
}

type user struct {
	Name string
}

func (u *user) String() string {
	return u.Name
}

func saveUser(user *user) {
	withTx(func() {
		fmt.Printf("保存用户 %s\n", user.Name)
	})
}

func withTx(fn func()) {
	fmt.Println("开启事务")
	fn()
	fmt.Println("结束事务")
}
