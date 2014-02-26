package main

import (
	"fmt"
	"time"
)

func DelayPrint(num int, wait int) {
	time.Sleep(time.Duration(wait) * time.Second)
	fmt.Println(num)
}

func main() {
	for i := 0; i < 6; i++ {
		go DelayPrint(i, 6 - i)
	}
	time.Sleep(10 * time.Second)
}
