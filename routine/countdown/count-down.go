/*
 * Copyright (c) 2023 fastbpmn
 * Project repo:https://github.com/fastbpmn/go-test
 */

package countdown

import "fmt"

func CountDown() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(<-c)
		}

		quit <- 0
	}()

	start(c, quit)
}

func start(c, quit chan int) {
	x := 5

	for {
		select {
		case c <- x:
			x -= 1
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
