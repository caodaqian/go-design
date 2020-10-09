// Package
// @Author      : caodaqian
// @CreateTime  : 2020-10-10 00:03:04
// @LastEditors : caodaqian
// @LastEditTime: 2020-10-10 00:09:28
// @Description : 工厂模式场景定义
//
package main

import "fmt"

type human interface {
	say()
	run()
}

type blackHuman struct {
}

func (b *blackHuman) say() {
	fmt.Println("黑人在说话")
}

func (b *blackHuman) run() {
	fmt.Println("黑人在跑")
}

type whiteHuman struct {
}

func (w *whiteHuman) say() {
	fmt.Println("白人在说话")
}

func (w *whiteHuman) run() {
	fmt.Println("白人在跑")
}

type yellowHuman struct {
}

func (y *yellowHuman) say() {
	fmt.Println("黄人在说话")
}

func (y *yellowHuman) run() {
	fmt.Println("黄人在跑")
}
