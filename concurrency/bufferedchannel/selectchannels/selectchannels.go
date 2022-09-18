package main

import "fmt"

func main() {
	chan1 := make(chan int, 1)
	chan1 <- 1

	select {
		case c1Val := <-chan1: //bu channeli dinler ve eger bu channelde veri olursa bunu bu degiskene atar
			fmt.Println(c1Val)
	}
}