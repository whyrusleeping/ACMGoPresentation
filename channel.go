package main

import "fmt"
import "time"
import "math/rand"

func ListenAndPrint(input chan int) {
	for {
		number := <-input
		fmt.Printf("Recieved %d!\n", number)
	}
}

func main() {
	mychan := make(chan int)
	go ListenAndPrint(mychan)
	for i := 0; i < 50; i++ {
		mychan<-rand.Intn(100)
		time.Sleep(time.Second)
	}
}
