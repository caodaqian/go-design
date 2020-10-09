// Package
// @Author      : caodaqian
// @CreateTime  : 2020-10-08 23:34:04
// @LastEditors : caodaqian
// @LastEditTime: 2020-10-09 12:07:46
// @Description : 单例模式实践1——饿汉模式

package main

import "fmt"

var intf *singleton

type singleton struct {
	Data string
}

func newSingle() *singleton {
	return intf
}

func init() {
	intf = &singleton{
		Data: "Hello World",
	}
}

func main() {
	test := newSingle()
	fmt.Println(test.Data)
}
