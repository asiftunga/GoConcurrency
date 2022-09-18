package main

import (
	"fmt"
	"sync"
)

//do not communicate by sharing memory instead share memory by communication
//pointerlari etrafa sacma channel olustur o channeli go routinelere ver

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1) //kac adet goroutine oldugu
	//channel initialization
	unbufferedChan := make(chan int)
	// unbufferedda es zamanli olarak bir data okuyabilir ve bir data yazabiliriz
	//bufferedlarda ise verdigimiz size a kadar bi locking olmuyor sonrasi lockingleniyor
	//channel declaration
	var unbufferedChan2 chan int

	fmt.Println(unbufferedChan)
	fmt.Println(unbufferedChan2) //nil

	//only can read from channel
	go func(unbufChan <-chan int) {
		//alt satirdaki islem blockingdir. Data gelene kadar bu islem bloklanir data geldikten sonra hayatina devam eder
		//blocks until data arrives
		value := <-unbufChan //channeldan bir deger okuyor
		fmt.Println(value)
		wg.Done() //goroutine islemi bitti
		//unbufChan <- 5 this is not work here
	}(unbufferedChan) //unbufferedChan parametre olarak goroutine anonim fonksiyonuna verilmistir
	unbufferedChan <- 1000 //mainde bu channela bir data basiliyor (channela ekleniyor gibi dusunebilirim)
	wg.Wait() //goroutineleri bekle demektir
}

//! guzel bir soru: go da nasil oluyorda channel pass by value ile parametre seklinde eklendiginde localde olusan channel (copy) bizim mainde bastigimiz degisken degerine ulasabiliyor? Channellar arkaplanda bir arraye unsafe pointer olarak point ediyor. Ordaki channel kopya olabilir ancak channellarin isaret ettigi adres ayni array adresidir