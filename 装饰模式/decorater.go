///////////////////////////////////////
// @Author      : caodaqian
// @CreateTime  : 2020-11-09 23:59:09
// @LastEditors : caodaqian
// @LastEditTime: 2020-11-10 00:05:53
// @Description : 装饰模式实现
///////////////////////////////////////
package main

import (
	"fmt"
)

type sugarSchoolReport struct {
	schoolReport
}

func (s *sugarSchoolReport) reportHighScore() {
	fmt.Println("最高分和我差不多")
}

func (s *sugarSchoolReport) reportSort() {
	fmt.Println("排名不是很后面")
}

func (s *sugarSchoolReport) report() {
	s.reportHighScore()
	s.schoolReport.report()
	s.reportSort()
}

func main() {
	report := schoolReport{60,80}
	report.report()
	sugarSchoolReport := sugarSchoolReport{report}
	sugarSchoolReport.report()
}