///////////////////////////////////////
// @Author      : caodaqian
// @CreateTime  : 2020-12-11 00:00:46
// @LastEditors : caodaqian
// @LastEditTime: 2020-12-13 13:23:14
// @Description : 桥梁模式实现
///////////////////////////////////////
package main

import "fmt"

type BCorp struct {
	product iProduct
}

func (b *BCorp) makeMoney() {
	b.product.beProducted()
	b.product.beSelled()
}

type iProduct interface {
	beProducted()
	beSelled()
}

type House struct{}

func (h *House)	beProducted() {
	fmt.Println("生产的房子造好了")
}

func (h *House) beSelled() {
	fmt.Println("生产的房子卖出去了")
}

type Ipad struct{}

func (i *Ipad)	beProducted() {
	fmt.Println("生产的ipad造好了")
}

func (i *Ipad) beSelled() {
	fmt.Println("生产的ipad卖出去了")
}

type HouseCorp struct {
	bc BCorp
}

func (hc *HouseCorp) makeMoney() {
	hc.bc.makeMoney()
	fmt.Println("房地产公司赚大钱了")
}

type CopyCorp struct{
	bc BCorp
}

func (cc *CopyCorp) init(pro iProduct) {
	cc.bc.product = pro
}

func (cc *CopyCorp) makeMoney() {
	cc.bc.makeMoney()
	fmt.Println("山寨公司赚大钱了")
}

func main() {
	hc := &HouseCorp{
		bc: BCorp{
			product: new(House),
		},
	}
	hc.makeMoney()

	ipad := new(Ipad)
	cc := new(CopyCorp)
	cc.init(ipad)
	cc.makeMoney()
}