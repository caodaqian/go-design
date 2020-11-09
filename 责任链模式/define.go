///////////////////////////////////////
// @Author      : caodaqian
// @CreateTime  : 2020-11-09 22:59:36
// @LastEditors : caodaqian
// @LastEditTime: 2020-11-09 23:32:37
// @Description : 责任链模式定义场景类
///////////////////////////////////////
package main

type iWomen interface {
	getType() int
	getRequest() string
}

type women struct {
	kind int
	request string
}

func (w *women) getType() int {
	return w.kind
}

func (w *women) getRequest() string {
	return w.request
}
