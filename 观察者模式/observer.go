///////////////////////////////////////
// @Author      : caodaqian
// @CreateTime  : 2020-11-23 22:25:27
// @LastEditors : caodaqian
// @LastEditTime: 2020-11-23 22:49:28
// @Description : 观察者模式实现
///////////////////////////////////////
package main

import "fmt"

type iObserver interface {
	dosometing(string)
	getName() string
}

type observer struct{
	name string
}

func (ob *observer) dosometing(con string) {
	fmt.Println(ob.getName()+"有事汇报：" + con)
}

func (ob *observer) getName() string {
	return ob.name
}

type iObservable interface {
	addObserver(ob iObserver)
	removeObserver(ob iObserver)
	notifyObservers(context string)
}

type observable struct {
	observers []iObserver
}

func (o *observable) addObserver(ob iObserver) {
	o.observers = append(o.observers, ob)
}

func (o *observable) removeObserver(name string) {
	for k, v := range o.observers {
		if v.getName() == name {
			o.observers = append(o.observers[0:k], o.observers[k+1:]...)
		}
	}
}

func (o *observable) notifyObservers(context string) {
	for _, v := range o.observers {
		v.dosometing(context)
	}
}

func (o *observable) haveLunch() {
	fmt.Println("吃午饭了")
	o.notifyObservers("吃午饭")
}

func (o *observable) haveFun() {
	fmt.Println("出去玩了")
	o.notifyObservers("出去玩")
}

func main() {
	// 小红 小黑 小蓝 去观察 小白
	red := &observer{"red"}
	black := &observer{"black"}
	blue := &observer{"blue"}

	white := &observable{}
	white.addObserver(red)
	white.addObserver(black)
	white.addObserver(blue)

	white.haveFun()

	// 此时小红离开了
	white.removeObserver(red.name)

	white.haveLunch()
}