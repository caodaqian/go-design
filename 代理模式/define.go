// Package
// @Author      : caodaqian
// @CreateTime  : 2020-10-15 22:38:07
// @LastEditors : caodaqian
// @LastEditTime: 2020-10-15 22:55:19
// @Description : 定义代理模式模板类
//
package main

import "fmt"

type iGamePlayer interface {
	login(user string, passwd string)
	killBoss()
	upgrade()
}

type gamePlayer struct {
	name string
}

func (g *gamePlayer) killBoss() {
	fmt.Printf("%s 在杀怪\n", g.name)
}

func (g *gamePlayer) upgrade() {
	fmt.Printf("%s 升级\n", g.name)
}

func (g *gamePlayer) login(user string, pd string) {
	fmt.Printf("登录名：%s，%s 用户已登录\n", user, g.name)
}

type gamePlayerProxy struct {
	myGamer iGamePlayer
}

func (gp *gamePlayerProxy) killBoss() {
	gp.myGamer.killBoss()
}

func (gp *gamePlayerProxy) login(user string, pd string) {
	gp.myGamer.login(user, pd)
}

func (gp *gamePlayerProxy) upgrade() {
	gp.myGamer.upgrade()
}
