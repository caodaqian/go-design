// Package
// @Author      : caodaqian
// @CreateTime  : 2020-10-13 23:48:28
// @LastEditors : caodaqian
// @LastEditTime: 2020-10-14 23:03:34
// @Description : 建造者模式
//
package main

import (
	"testing"
)

func Test_director_getBCar(t *testing.T) {
	dir := new(director)
	dir.sequence = []string{"start", "alert", "stop"}
	dir.getACar().run()
	dir.sequence = []string{"alert", "start", "stop"}
	dir.getBCar().run()
}
