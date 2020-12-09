///////////////////////////////////////
// @Author      : caodaqian
// @CreateTime  : 2020-12-04 00:08:05
// @LastEditors : caodaqian
// @LastEditTime: 2020-12-09 22:51:15
// @Description : 解释器模式实现
///////////////////////////////////////
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Calculator struct {
	expStr iExpression
}

func (c *Calculator) calculate(exp string) {
	stack := []iExpression{}

	for i := 0; i < len(exp); i++ {
		switch exp[i] {
		case '+':
			left := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			i++
			right := &VarExpression{string(exp[i])}
			stack = append(stack, &AddExpression{
				mse: SymbolExpression{
					left:  left,
					right: right,
				},
			})
		case '-':
			left := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			i++
			right := &VarExpression{string(exp[i])}
			stack = append(stack, &SubExpression{
				mse: SymbolExpression{
					left:  left,
					right: right,
				},
			})
		default:
			stack = append(stack, &VarExpression{string(exp[i])})
		}
	}
	c.expStr = stack[len(stack)-1]
}

func (c *Calculator) run(m map[string]int) int {
	return c.expStr.interpreter(m)
}

type iExpression interface {
	interpreter(m map[string]int) int
}

type VarExpression struct {
	key string
}

func (ve *VarExpression) interpreter(m map[string]int) int {
	return m[ve.key]
}

type SymbolExpression struct {
	left  iExpression
	right iExpression
}

type AddExpression struct {
	mse SymbolExpression
}

func (ae *AddExpression) interpreter(m map[string]int) int {
	return ae.mse.left.interpreter(m) + ae.mse.right.interpreter(m)
}

type SubExpression struct {
	mse SymbolExpression
}

func (se *SubExpression) interpreter(m map[string]int) int {
	return se.mse.left.interpreter(m) - se.mse.right.interpreter(m)
}

func buildMap(exp string) *map[string]int {
	m := make(map[string]int, len(exp)/2+1)
	input := bufio.NewScanner(os.Stdin)
	for _, v := range exp {
		if v != '+' && v != '-' {
			if _, ok := m[string(v)]; !ok {
				fmt.Printf("请输入 %s 的值:", string(v))
				input.Scan()
				m[string(v)], _ = strconv.Atoi(input.Text())
			}
		}
	}
	return &m
}

func main() {
	exp := "a+b-c+d"
	cal := new(Calculator)
	m := buildMap(exp)
	cal.calculate(exp)
	fmt.Printf("exp=%s, result=%d\n", exp, cal.run(*m))
}