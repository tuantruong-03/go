package main

import (
	"fmt"
)

func doSth() {
	forever := make(chan bool)
	go func() {
		fmt.Println("do sth")
		for {

		}
	}()

	<-forever
}
