package create

import (
	"fmt"
	"sync"
)

type Config struct {
	Key, Val string
}

var (
	once sync.Once
	// 全局单实例
	instance *Config
)

func GetConf() *Config {
	once.Do(func() {
		fmt.Println("init")
		instance = new(Config)
	})

	return instance
}
