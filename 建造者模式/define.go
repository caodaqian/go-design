// Package
// @Author      : caodaqian
// @CreateTime  : 2020-10-12 23:54:16
// @LastEditors : caodaqian
// @LastEditTime: 2020-10-14 22:47:40
// @Description : 定义建造者模式的场景类
//
package main

import "fmt"

type iCar interface {
	start()
	stop()
	alert()
	set([]string)
}

type car struct {
	iCar
	mySeqence []string
}

func (c *car) run() {
	for _, v := range c.mySeqence {
		switch v {
		case "start":
			c.start()
		case "stop":
			c.stop()
		case "alert":
			c.alert()
		default:
			panic("unknown action")
		}
	}
}

type aCar struct {
	car
}

func (a *aCar) start() {
	fmt.Println("a 车启动")
}

func (a *aCar) stop() {
	fmt.Println("a 车停下")
}

func (a *aCar) alert() {
	fmt.Println("a 车鸣笛")
}

type bCar struct {
	car
}

func (b *bCar) start() {
	fmt.Println("b 车启动")
}

func (b *bCar) stop() {
	fmt.Println("b 车停下")
}

func (b *bCar) alert() {
	fmt.Println("b 车鸣笛")
}
