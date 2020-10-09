// Package
// @Author      : caodaqian
// @CreateTime  : 2020-10-10 00:15:06
// @LastEditors : caodaqian
// @LastEditTime: 2020-10-10 00:17:40
// @Description : 简单工厂模式
//
package main

func buildHuman(x int) human {
	switch x {
	case 1:
		return &whiteHuman{}
	case 2:
		return &blackHuman{}
	case 3:
		return &yellowHuman{}
	default:
		return nil
	}
}

func main() {
	people := buildHuman(1)
	people.run()
	people.say()
}
