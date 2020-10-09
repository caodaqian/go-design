// Package
// @Author      : caodaqian
// @CreateTime  : 2020-10-10 00:00:01
// @LastEditors : caodaqian
// @LastEditTime: 2020-10-10 00:12:27
// @Description : 抽象工厂模式
//
package main

type abstractFactory interface {
	buildHuman() human
}

type whiteFactory struct {
}

func (wf *whiteFactory) buildHuman() human {
	return &whiteHuman{}
}

type blackFactory struct {
}

func (bf *blackFactory) buildHuman() human {
	return &blackHuman{}
}

type yellowFactory struct {
}

func (yf *yellowFactory) buildHuman() human {
	return &yellowHuman{}
}

func main() {
	fact := new(yellowFactory)
	people := fact.buildHuman()
	people.run()
	people.say()
}
