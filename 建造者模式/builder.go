// Package
// @Author      : caodaqian
// @CreateTime  : 2020-10-13 23:48:28
// @LastEditors : caodaqian
// @LastEditTime: 2020-10-14 22:56:18
// @Description : 建造者模式
//
package main

type iCarBuilder interface {
	setSeqence([]string)
	getCar() *car
}

type aCarBuilder struct {
	myCar aCar
}

func (acb *aCarBuilder) getCar() *car {
	return &acb.myCar.car
}

func (acb *aCarBuilder) setSeqence(seq []string) {
	acb.myCar.mySeqence = seq
}

type bCarBuilder struct {
	myCar bCar
}

func (bcb *bCarBuilder) getCar() *car {
	return &bcb.myCar.car
}

func (bcb *bCarBuilder) setSeqence(seq []string) {
	bcb.myCar.mySeqence = seq
}

type director struct {
	sequence    []string
	carABuilder aCarBuilder
	carBBuilder bCarBuilder
}

func (d *director) getACar() *car {
	d.sequence = []string{}
	d.sequence = append(d.sequence, "start", "alert", "stop")
	d.carABuilder.setSeqence(d.sequence)
	return d.carABuilder.getCar()
}

func (d *director) getBCar() *car {
	d.sequence = []string{}
	d.sequence = append(d.sequence, "alert", "start", "stop")
	d.carBBuilder.setSeqence(d.sequence)
	return d.carBBuilder.getCar()
}

func main() {
	dir := new(director)
	dir.getACar().run()
	dir.getBCar().run()
}
