///////////////////////////////////////
// @Author      : caodaqian
// @CreateTime  : 2020-12-09 23:33:10
// @LastEditors : caodaqian
// @LastEditTime: 2020-12-10 23:58:34
// @Description : 享元模式实现
///////////////////////////////////////
package main

import (
	"fmt"
	"strconv"
)

type SignInfo struct {
	id          string
	location    string
	subject     string
	postAddress string
}

type SignInfoPool struct {
	key  string
	info SignInfo
}

type SignInfoFactory struct {
	pool map[string]*SignInfoPool
}

func (sf *SignInfoFactory) getSignInfo(key string) *SignInfoPool {
	val, ok := sf.pool[key]
	if !ok {
		// 没有则创建
		fmt.Println("创建一个可共享对象,key:"+key)
		res := &SignInfoPool{
			key: key,
		}
		sf.pool[key] = res
		return res
	}
	return val
}

func main() {
	factory := new(SignInfoFactory)
	factory.pool = make(map[string]*SignInfoPool, 40)
	for i := 0; i < 10; i++ {
		s := "科目:" + strconv.Itoa(i)
		for j := 0; j < 20; j++ {
			ss := s + "考试地点:" + strconv.Itoa(j)
			factory.getSignInfo(ss)
		}
	}
	fmt.Printf("%+v", factory.pool)
}
