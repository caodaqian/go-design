// Package
// @Author      : caodaqian
// @CreateTime  : 2020-10-13 23:48:28
// @LastEditors : caodaqian
// @LastEditTime: 2020-10-15 22:11:22
// @Description : 建造者模式
//
package main

type iCarBuilder interface {
	setSeqence([]string)
	getCar() *car
}

type aCarBuilder struct {
	myACar car
}

func (acb *aCarBuilder) getCar() *car {
	acb.myACar.iCar = new(aCar)
	return &acb.myACar
}

func (acb *aCarBuilder) setSeqence(seq []string) {
	acb.myACar.mySeqence = seq
}

type bCarBuilder struct {
	myBCar car
}

func (bcb *bCarBuilder) getCar() *car {
	bcb.myBCar.iCar = new(bCar)
	return &bcb.myBCar
}

func (bcb *bCarBuilder) setSeqence(seq []string) {
	bcb.myBCar.mySeqence = seq
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
