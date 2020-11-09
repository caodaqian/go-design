///////////////////////////////////////
// @Author      : caodaqian
// @CreateTime  : 2020-11-10 00:09:45
// @LastEditors : caodaqian
// @LastEditTime: 2020-11-10 00:16:41
// @Description : 策略模式实现
///////////////////////////////////////
package main

import "fmt"

type iIdea interface {
	operate()
}

type ideaA struct {}

func (a *ideaA) operate() {
	fmt.Println("找乔国老")
}

type ideaB struct {}

func (b *ideaB) operate() {
	fmt.Println("找吴国太")
}

type ideaC struct {}

func (c *ideaC) operate() {
	fmt.Println("找孙夫人")
}

type context struct {
	midea iIdea
}

func (c *context) init(idea iIdea) {
	c.midea = idea
}

func main() {
	c := new(context)

	c.init(new(ideaA))
	c.midea.operate()
	c.init(new(ideaB))
	c.midea.operate()
	c.init(new(ideaC))
	c.midea.operate()
}