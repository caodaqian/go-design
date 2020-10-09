// Package
// @Author      : caodaqian
// @CreateTime  : 2020-10-09 12:10:28
// @LastEditors : caodaqian
// @LastEditTime: 2020-10-09 12:10:49
// @Description : 单例模式实践1——懒汉模式

package main

import (
	"fmt"
	"sync"
)

var intf *singleton
var mut sync.Mutex

type singleton struct {
	Data string
}

func newSingle() *singleton {
	// 锁 + 双重判断
	// @import 为了防止编译器优化，使用 go build 时应加上 -gcflags "-N"
	// fixme 如何实现 go 版本 volatile 语义？
	if intf == nil {
		mut.Lock()
		if intf == nil {
			intf = &singleton{
				Data: "Hello World",
			}
		}
		mut.Unlock()
	}
	return intf
}

func main() {
	test := newSingle()
	fmt.Println(test.Data)
}
