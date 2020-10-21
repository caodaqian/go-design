// Package
// @Author      : caodaqian
// @CreateTime  : 2020-10-21 23:07:17
// @LastEditors : caodaqian
// @LastEditTime: 2020-10-21 23:59:02
// @Description : 命令模式示例
//
package main

type command interface {
	execute()
}

type baseCommand struct{
	mPageGroup pageGroup
	mRequirementGroup requirementGroup
}

type addRequirementCommand struct {
	baseCommand
}

func (arc *addRequirementCommand) execute() {
	arc.baseCommand.mRequirementGroup.find()
	arc.baseCommand.mRequirementGroup.add()
}

type deletePageCommand struct {
	baseCommand
}

func (dpc *deletePageCommand) execute() {
	dpc.baseCommand.mPageGroup.find()
	dpc.baseCommand.mPageGroup.del()
}

type invoker struct {
	saveCommand command
}

func (i *invoker) setCommand(bc command) {
	if bc == nil {
		panic("empty command")
	}
	i.saveCommand = bc
}

func (i *invoker) action() {
	i.saveCommand.execute()
}

func main() {
	manager := &invoker{}

	cmd1 := &addRequirementCommand{}
	manager.setCommand(cmd1)
	manager.action()

	cmd2 := &deletePageCommand{}
	manager.setCommand(cmd2)
	manager.action()
}