/*
 * Copyright (c) 2023 fastbpmn
 * Project repo:https://github.com/fastbpmn/go-test
 */

package main

import (
	"github.com/fastbpmn/go-study/date-time/dt"
	"github.com/fastbpmn/go-test/routine/countdown"
)

func main() {
	countdown.CountDown()
	dt.Day()
	dt.PrintDT()
}
