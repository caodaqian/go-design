///////////////////////////////////////
// @Author      : caodaqian
// @CreateTime  : 2020-11-30 23:29:27
// @LastEditors : caodaqian
// @LastEditTime: 2020-12-01 00:29:11
// @Description : 访问者模式实现
///////////////////////////////////////
package main

import (
	"fmt"
)

type iEmployee interface {
	getOtherInfo() string
	report() string
}

type baseEmployee struct {
	name   string
	salary int
}

type commonEmployee struct {
	baseEmployee
	job string
}

func (ce *commonEmployee) report() string {
	return fmt.Sprintf("name:%s\tsalary:%d\t%s", ce.name, ce.salary, ce.getOtherInfo())
}

func (ce *commonEmployee) getOtherInfo() string {
	return "job:" + ce.job
}

type mananger struct {
	baseEmployee
	performance string
}

func (m *mananger) getOtherInfo() string {
	return "performance:" + m.performance
}

func (m *mananger) report() string {
	return fmt.Sprintf("name:%s\tsalary:%d\t%s", m.name, m.salary, m.getOtherInfo())
}

type iVisitor interface {
	visit(iEmployee)
}

type visitor struct {
}

func (v *visitor) visit(ie iEmployee) {
	switch ie.(type) {
	case *commonEmployee:
		fmt.Println(ie.report())
	case *mananger:
		fmt.Println(ie.report())
	default:
		panic("unknown employee type")
	}
}

func main() {
	zhangsan := &commonEmployee{
		baseEmployee: baseEmployee{
			name:   "zhangsan",
			salary: 1184,
		},
		job: "开发",
	}
	lisi := &mananger{
		baseEmployee: baseEmployee{
			name:   "lisi",
			salary: 1184000,
		},
		performance: "项目进行中",
	}

	vis := new(visitor)
	vis.visit(zhangsan)
	vis.visit(lisi)
}