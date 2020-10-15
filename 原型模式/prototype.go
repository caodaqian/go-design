// Package
// @Author      : caodaqian
// @CreateTime  : 2020-10-15 23:48:43
// @LastEditors : caodaqian
// @LastEditTime: 2020-10-15 23:54:20
// @Description : 原型模式
//
package main

import "fmt"

type example struct {
	content string
}

func (e *example) clone() *example {
	return &example{
		content: e.content,
	}
}

func main() {
	e1 := example{"test"}
	e2 := e1.clone()
	e3 := e2.clone()
	fmt.Printf("e1: %+v, point:%p\n", e1, &e1)
	fmt.Printf("e2: %+v, point:%p\n", *e2, e2)
	fmt.Printf("e3: %+v, point:%p\n", *e3, e3)
}
