///////////////////////////////////////
// @Author      : caodaqian
// @CreateTime  : 2020-11-22 23:18:13
// @LastEditors : caodaqian
// @LastEditTime: 2020-11-23 22:18:31
// @Description : 策略模式实现
///////////////////////////////////////
package main

import "fmt"

type iCorp interface {
	// 获取每个员工的信息接口
	getInfo() string
}

// 代表打工仔
type leaf struct {
	name   string
	title  string
	salary int
}

func (l *leaf) getInfo() string {
	return fmt.Sprintf("该员工姓名：%s,职位:%s，薪水：%d", l.name, l.title, l.salary)
}

type iBranch interface {
	addChild(c iCorp)
	getChild() []iCorp
}

// 代表领导
type branch struct {
	name   string
	title  string
	salary int
	employ []iCorp
}

func (b *branch) addChild(c iCorp) {
	b.employ = append(b.employ, c)
}

func (b *branch) getChild() []iCorp {
	return b.employ
}

func (b *branch) getInfo() string {
	return fmt.Sprintf("该领导姓名：%s,职位:%s，薪水：%d", b.name, b.title, b.salary)
}

// --------------------

func main() {
	leaf1 := &leaf{
		name:   "leaf1",
		title:  "leaf",
		salary: 111,
	}
	leaf2 := &leaf{
		name:   "leaf2",
		title:  "leaf",
		salary: 111,
	}
	leader1 := &branch{
		name:   "leader1",
		title:  "leader",
		salary: 1111,
		employ: []iCorp{leaf1, leaf2},
	}
	leaf3 := &leaf{
		name:   "leaf3",
		title:  "leaf",
		salary: 111,
	}
	leader2 := &branch{
		name:   "leader2",
		title:  "leader",
		salary: 1111,
		employ: []iCorp{leaf3},
	}
	ceo := &branch{
		name:   "ceo",
		title:  "CEO",
		salary: 11110,
		employ: []iCorp{leader1, leader2},
	}

	queue := []iCorp{ceo}
	for len(queue) != 0 {
		node := queue[0]
		if v, ok := node.(iBranch); ok {
			queue = append(queue, v.getChild()...)
		}
		queue = queue[1:]
		fmt.Println(node.getInfo())
	}
}
