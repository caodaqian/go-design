///////////////////////////////////////
// @Author      : caodaqian
// @CreateTime  : 2020-11-09 23:52:42
// @LastEditors : caodaqian
// @LastEditTime: 2020-11-09 23:57:47
// @Description : 装饰模式定义场景类
///////////////////////////////////////
package main

import "fmt"

type schoolReport struct {
	chinese int
	math    int
}

func (s *schoolReport) report() {
	fmt.Printf("成绩: 语文%d 数学%d\n", s.chinese, s.math)
}