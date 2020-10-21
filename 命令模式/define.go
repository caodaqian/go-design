// Package
// @Author      : caodaqian
// @CreateTime  : 2020-10-21 23:09:49
// @LastEditors : caodaqian
// @LastEditTime: 2020-10-21 23:24:56
// @Description : 命令模式定义场景类
//
package main

import "fmt"

type group interface {
	find()
	add()
	del()
}

type requirementGroup struct {}

func (r *requirementGroup) find() {
	fmt.Println("找到需求组")
}

func (r *requirementGroup) add() {
	fmt.Println("添加一个需求")
}

func (r *requirementGroup) del() {
	fmt.Println("删除一个需求")
}

type pageGroup struct {}

func (p *pageGroup) find() {
	fmt.Println("找到美工组")
}

func (p *pageGroup) add() {
	fmt.Println("添加一个页面")
}

func (p *pageGroup) del() {
	fmt.Println("删除一个页面")
}