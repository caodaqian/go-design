///////////////////////////////////////
// @Author      : caodaqian
// @CreateTime  : 2020-12-01 23:40:28
// @LastEditors : caodaqian
// @LastEditTime: 2020-12-03 23:26:11
// @Description : 状态模式实现
///////////////////////////////////////
package main

import "fmt"

type iLiftState interface {
	open() iLiftState
	close() iLiftState
	move() iLiftState
	stop() iLiftState
}

const (
	OpenState  = 1
	CloseState = 2
	MoveState  = 3
	StopState  = 4
)

type OpenningState struct {}

func (os *OpenningState) open() iLiftState {
	fmt.Println("电梯门打开...")
	return os
}

func (os *OpenningState) close() iLiftState {
	fmt.Println("电梯门切换到关闭")
	return new(ClosingState)
}

func (os *OpenningState) move() iLiftState {
	fmt.Println("电梯无法移动")
	return os
}

func (os *OpenningState) stop() iLiftState {
	fmt.Println("电梯已经停止")
	return os
}

type ClosingState struct {}

func (cs *ClosingState) open() iLiftState {
	fmt.Println("电梯门切换到开启")
	return new(OpenningState)
}

func (cs *ClosingState) close() iLiftState {
	fmt.Println("电梯门正在关闭")
	return cs
}

func (cs *ClosingState) move() iLiftState {
	fmt.Println("电梯即将启动")
	return cs // 省略一个移动状态，cs代替一下
}

func (cs *ClosingState) stop() iLiftState {
	fmt.Println("电梯即将停止")
	return cs // 省略一个停止状态，cs代替一下
}

type lift struct {
	context iLiftState
}

func main() {
	l := &lift{
		new(OpenningState),
	}

	l.context = l.context.open()
	l.context = l.context.move()
	l.context = l.context.stop()
	l.context = l.context.close()

	l.context = l.context.close()
	l.context = l.context.move()
	l.context = l.context.stop()
	l.context = l.context.open()

	l.context.open()
}
