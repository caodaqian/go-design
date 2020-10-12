// Package
// @Author      : caodaqian
// @CreateTime  : 2020-10-12 23:43:40
// @LastEditors : caodaqian
// @LastEditTime: 2020-10-12 23:49:27
// @Description : 模板方法模式的场景类与示例
//
package main

import "fmt"

type car interface {
	start()
	stop()
}

func run(c car) {
	c.start()
	c.stop()
}

type carA struct {
}

func (c *carA) start() {
	fmt.Println("car A is starting")
}

func (c *carA) stop() {
	fmt.Println("car A is stopping")
}

type carB struct {
}

func (c *carB) start() {
	fmt.Println("car B is starting")
}

func (c *carB) stop() {
	fmt.Println("car B is stopping")
}

func main() {
	a := new(carA)
	run(a)
	b := new(carB)
	run(b)
}
