package main

import (
	"fmt"
	"sync"
)

type RaceTest struct {
	Val int
}

//10000 tane olusturdugumuz halde 10000 tane olusmadi bunu cozmenin yolu mutex (WaitGroup) dir (racecondition.go da var)

func main() {
	raceTest := &RaceTest{}

	wg := &sync.WaitGroup{}
	wg.Add(10000)

	for i:=0; i<10000; i++ {
		go increment(raceTest, wg)
	}

	wg.Wait()

	fmt.Println(raceTest)
}

func increment(rt *RaceTest, wg *sync.WaitGroup) {
	rt.Val += 1
	wg.Done()
}