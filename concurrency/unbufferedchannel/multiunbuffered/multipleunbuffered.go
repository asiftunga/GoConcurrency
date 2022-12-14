package main

import (
	"fmt"
	"time"
)

func main() {
	//channel initialization
	unbufferedChan := make(chan int)

	//reader goroutine
	go func(unbufChan chan int) {
		value := <-unbufChan
		fmt.Println(value)
	}(unbufferedChan)

	//writer goroutine
	go func(unbufChan chan int) {
		unbufChan <- 1
	}(unbufferedChan)


	//now scheduler have time to schedule, still, this might not work and not the best solution
	//kisaca alttaki islem main routine bitmeden digerlerine zaman kalsin diye yapilmis bir olay
	//mutex kullanabilirdik
	time.Sleep(time.Second)

	fmt.Println("hello channels")

}