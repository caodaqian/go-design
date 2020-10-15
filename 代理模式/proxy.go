// Package
// @Author      : caodaqian
// @CreateTime  : 2020-10-15 22:50:12
// @LastEditors : caodaqian
// @LastEditTime: 2020-10-15 22:55:55
// @Description : 代理模式
//
package main

import (
	"fmt"
	"time"
)

func baseProxyPattern() {
	realGamer := gamePlayer{
		name: "realname",
	}
	proxyGamer := gamePlayerProxy{
		myGamer: &realGamer,
	}
	fmt.Println("开始时间" + time.Now().Format("2016-01-02 15:04:05"))
	proxyGamer.login("zhangsan", "123213")
	proxyGamer.killBoss()
	proxyGamer.upgrade()
	fmt.Println("结束时间" + time.Now().Format("2016-01-02 15:04:05"))
}

func main() {
	baseProxyPattern()
}
