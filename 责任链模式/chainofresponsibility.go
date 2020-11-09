///////////////////////////////////////
// @Author      : caodaqian
// @CreateTime  : 2020-11-09 23:35:29
// @LastEditors : caodaqian
// @LastEditTime: 2020-11-09 23:51:29
// @Description : 责任链模式实现
///////////////////////////////////////
package main

import "fmt"

type iHandler interface {
	response(iWomen)
}

type handler struct {
	mLevel int
	nextHandler *handler
}

func (h *handler) response(i iWomen) {
	if i.getType() == h.mLevel {
		fmt.Printf("Level:%d handler response OK with kind:%d woman request:%s\n", h.mLevel, i.getType(), i.getRequest())
	} else if h.nextHandler != nil {
		h.nextHandler.response(i)
	} else {
		fmt.Printf("Level:%d handler response NOT with kind:%d woman request:%s\n", h.mLevel, i.getType(), i.getRequest())
	}
}

func main() {
	h3 := handler{3, nil}
	h2 := handler{2, &h3}
	h1 := handler{1, &h2}

	w1 := &women{1,"1号请求"}
	w2 := &women{2,"2号请求"}
	w3 := &women{3,"3号请求"}
	w4 := &women{4,"4号请求"}

	h1.response(w1)
	h1.response(w2)
	h1.response(w3)
	h1.response(w4)
}