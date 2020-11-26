///////////////////////////////////////
// @Author      : caodaqian
// @CreateTime  : 2020-11-26 23:47:17
// @LastEditors : caodaqian
// @LastEditTime: 2020-11-26 23:53:20
// @Description : 备忘录模式
///////////////////////////////////////
package main

import "fmt"

type test struct {
	status string
}

type iManager interface {
	pushBackup(test)
	popBackup() test
}

type manager struct {
	backups []test
}

func (m *manager) pushBackup(t test) {
	m.backups = append(m.backups, t)
}

func (m *manager) popBackup() test {
	defer func() {
		m.backups = m.backups[:len(m.backups)-1]
	}()
	return m.backups[len(m.backups)-1]
}

func main() {
	t1 := test{"happy"}
	m := manager{}

	fmt.Println("一开始状态", t1)
	m.pushBackup(t1)
	t1.status = "sad"
	fmt.Println("现在状态变了", t1)
	t1 = m.popBackup()
	fmt.Println("回退到上一个状态", t1)
}
