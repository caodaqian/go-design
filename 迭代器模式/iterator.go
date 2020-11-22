///////////////////////////////////////
// @Author      : caodaqian
// @CreateTime  : 2020-11-15 23:43:15
// @LastEditors : caodaqian
// @LastEditTime: 2020-11-16 00:15:13
// @Description : 迭代器模式实现
///////////////////////////////////////
package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type baseProject struct {
	name string
	cost int
}

func (bp *baseProject) getProjectInfo() string {
	return fmt.Sprintf("项目名称:%s\t项目花费:%d", bp.name, bp.cost)
}

type projects struct {
	data []*baseProject
	current int
}

func (p *projects) add(item *baseProject) {
	p.data = append(p.data, item)
}

func (p *projects) hasNext() bool {
	fmt.Println(p.current)
	return len(p.data) > p.current
}

func (p *projects) next() *baseProject {
	if p.current < len(p.data) {
		p.current++
		return p.data[p.current-1]
	}
	return nil
}

func main() {
	pros := new(projects)
	pros.current = -1
	for i:=0;i<10;i++ {
		bp := &baseProject{"test"+strconv.Itoa(i), rand.Intn(100)}
		pros.add(bp)
	}
	for item := pros.next();pros.hasNext();item = pros.next() {
		fmt.Println(item.getProjectInfo())
	}
}