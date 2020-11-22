///////////////////////////////////////
// @Author      : caodaqian
// @CreateTime  : 2020-11-15 23:06:29
// @LastEditors : caodaqian
// @LastEditTime: 2020-11-15 23:28:03
// @Description : 适配器模式实现
///////////////////////////////////////
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type iUser interface {
	getName() string
	getHome() string
	getPhone() string
}

type userinfo struct {
	name  string
	home  string
	phone string
}

func (u *userinfo) getName() string {
	return u.name
}

func (u *userinfo) getHome() string {
	return u.home
}

func (u *userinfo) getPhone() string {
	return u.phone
}

type iOuterUser interface {
	getName() map[string]string
	getHome() map[string]string
	getPhone() map[string]string
}

type outerUser struct {
	userinfo map[string]string
}

func (ou *outerUser) getName() map[string]string {
	return ou.userinfo
}

func (ou *outerUser) getHome() map[string]string {
	return ou.userinfo
}

func (ou *outerUser) getPhone() map[string]string {
	return ou.userinfo
}

// adapter
type outerUserInfo struct {
	outerUser
}

func (oui *outerUserInfo) getName() string {
	return oui.userinfo["name"]
}

func (oui *outerUserInfo) getHome() string {
	return oui.userinfo["home"]
}

func (oui *outerUserInfo) getPhone() string {
	return oui.userinfo["phone"]
}

func main() {
	// normal
	rand.Seed(time.Now().Unix())
	for i := 0; i < 5; i++ {
		localUser := userinfo{
			name:  "test name",
			home:  "test home",
			phone: strconv.Itoa(rand.Intn(9999999)),
		}
		fmt.Printf("本地的员工名：%s\t员工家庭地址：%s\t员工电话：%s\n", localUser.getName(), localUser.getHome(), localUser.getPhone())
	}
	for i := 0; i < 5; i++ {
		outeruser := outerUserInfo{
			outerUser: outerUser{
				userinfo: map[string]string{
					"name":  "测试名",
					"home":  "测试家庭地址",
					"phone": strconv.Itoa(rand.Intn(9999999)),
				},
			},
		}
		fmt.Printf("外包的员工名：%s\t员工家庭地址：%s\t员工电话：%s\n", outeruser.getName(), outeruser.getHome(), outeruser.getPhone())
	}
}
