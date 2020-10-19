// Package
// @Author      : caodaqian
// @CreateTime  : 2020-10-15 23:56:19
// @LastEditors : caodaqian
// @LastEditTime: 2020-10-20 00:18:03
// @Description : 中介者模式 （go 不能实现抽象类，所以导致类关系有些模糊）
//
package main

import (
	"fmt"
	"math/rand"
)

type iMediator interface {
	execute(event string, obj ...interface{})
}

type mediator struct {
	mPurchase *purchase
	mSale     *sale
	mStock    *stock
}

func (m *mediator) execute(event string, obj ...interface{}) {
	switch event {
	case "purchase.buy":
		m.buy(obj[0].(int))
	case "sale.sell":
		m.sell(obj[0].(int))
	case "sale.offsell":
		m.offsell()
	case "stock.clear":
		m.clearStock()
	}
}

func (m *mediator) buy(num int) {
	saleStatus := m.mSale.getSaleStatus()
	if saleStatus > 80 {
		// 销售状况良好
		fmt.Printf("采购 %d 台", num)
		m.mStock.increase(num)
	} else {
		// 销售状况不好
		fmt.Printf("采购 %d 台", num/2)
		m.mStock.increase(num / 2)
	}
}

func (m *mediator) sell(num int) {
	if count := m.mStock.getStockCount(); count < num {
		// 库存不够销售
		m.mPurchase.buy(num - count)
	}
	m.mStock.decrease(num)
}

func (m *mediator) offsell() {
	fmt.Printf("折价销售 %d 台", m.mStock.getStockCount())
}

func (m *mediator) clearStock() {
	m.mSale.offsell()
	m.mPurchase.refuseBuy()
}

type purchase struct {
	mm *mediator
}

func (p *purchase) buy(num int) {
	p.mm.execute("purchase.buy", num)
}

func (p *purchase) refuseBuy() {
	fmt.Println("不再采购")
}

type stock struct {
	mm    *mediator
	count int
}

func (s *stock) increase(num int) {
	s.count += num
	fmt.Printf("库存增长到 %d \n", s.count)
}

func (s *stock) decrease(num int) {
	s.count -= num
	fmt.Printf("库存减少到 %d \n", s.count)
}

func (s *stock) getStockCount() int {
	return s.count
}

func (s *stock) clearStock() {
	fmt.Printf("清理存货 %d \n", s.count)
	s.mm.execute("stock.clear")
}

type sale struct {
	mm *mediator
}

func (s *sale) sell(num int) {
	s.mm.execute("sale.sell", num)
	fmt.Printf("卖出 %d \n", num)
}

func (s *sale) getSaleStatus() int {
	mock := rand.Int()
	return mock
}

func (s *sale) offsell() {
	s.mm.execute("sale.offsell")
}

func main() {
	med := new(mediator)
	med.mPurchase = &purchase{med}
	med.mSale = &sale{med}
	med.mStock = &stock{med, 50}

	med.mPurchase.buy(100)
	med.mSale.sell(10)
	med.mStock.clearStock()
}
