package main

import "fmt"

func main() {
	//channel initialization
	bufferedChan := make(chan int, 5) //buffered channelin farki size verilebilmesidir. Unbufferedda hatirlanirsa sadece bir kere okuma ve yazma yapilabilir (ayni anda)
	//bu sizedan buyuk bir write islemi yapilirsa (ornegin alttakinden bir tane daha fazla) deadlock olayi gerceklesir

	bufferedChan <-1
	bufferedChan <-2
	bufferedChan <-3
	bufferedChan <-4
	bufferedChan <-5

	fmt.Println(<-bufferedChan)
	fmt.Println(<-bufferedChan)
	fmt.Println(<-bufferedChan)
	fmt.Println(<-bufferedChan)
	fmt.Println(<-bufferedChan)
}