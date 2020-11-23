///////////////////////////////////////
// @Author      : caodaqian
// @CreateTime  : 2020-11-23 23:08:42
// @LastEditors : caodaqian
// @LastEditTime: 2020-11-23 23:18:20
// @Description : 门面模式实现
///////////////////////////////////////
package main

import "fmt"

type  iLetterProcess interface {
	writeContent(string)
	fillEnvelope(string)
	letterIntoEnvelope()
	sendLetter()
}

type letterProcesser struct {}

func (l *letterProcesser) writeContent(con string) {
	fmt.Println("写信的内容："+con)
}

func (l *letterProcesser) fillEnvelope(con string)  {
	fmt.Println("写信封："+con)
}

func (l *letterProcesser) letterIntoEnvelope() {
	fmt.Println("把信装进信封")
}

func (l *letterProcesser) sendLetter() {
	fmt.Println("发信")
}

type modenPostOffice struct {
	ml letterProcesser
}

func (m *modenPostOffice) sendLetter(con string, title string) {
	fmt.Println("现代邮寄信方式")
	m.ml.writeContent(con)
	m.ml.fillEnvelope(title)
	m.ml.letterIntoEnvelope()
	m.ml.sendLetter()
}

func main() {
	moden := new(modenPostOffice)
	moden.sendLetter("我的信的内容", "我要邮寄到那里去")
}