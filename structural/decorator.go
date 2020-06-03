package structural

import (
	"log"
	"strings"
)

type StringHandler func(string) string

func LogDecorate(fn StringHandler) StringHandler {
	return func(s string) string {
		log.Println("Starting the execution with the String", s)

		result := fn(s)

		log.Println("Execution is completed with the result", result)

		return result
	}
}

func ReplaceDecorate(fn StringHandler, oldStr, replaceStr string) StringHandler {
	return func(s string) string {
		newStr := strings.Replace(s, oldStr, replaceStr, -1)

		// 调用原函数
		result := fn(newStr)

		return result
	}
}
