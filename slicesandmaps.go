package main

import (
	"fmt"
)


var map1 map[int]string

func main(){
 	println("================array ve slice tanimlamalari=======================")
	slice1 := make([]int, 0, 3)
	array1 := make([]int, 3)
	fmt.Println("slice: ",slice1,"array: ",array1," array len i: ",len(array1)," slice len i: ",len(slice1), " slice cap i: ",cap(slice1))
	slice1 = append(slice1,13)
	slice1 = append(slice1,23)
	slice1 = append(slice1,43)
	slice1 = append(slice1,63)
	fmt.Println("slice: ",slice1,"array: ",array1," array len i: ",len(array1)," slice len i: ",len(slice1), " slice cap i: ",cap(slice1))
	// kapasite fazladan bir eleman eklendiginde direkt olarak belirlenen bir kapasitenin iki katina cikar. Go slice calisma mantigi bu sekildedir. Baslangic olarak herhangi bir kapasite verilmez ise default kapasite 2 dir sonrasinda her gecildiginde 2 katina cikar
	fmt.Println("====================== map tanimlamalari =========================")
	map1 = make(map[int]string) //bu map fonksiyon disinda tanimlandi
	map2 := make(map[int]int)
	map3 := make(map[string]int, 3)
	map4 := map[int]string{
		1923: "TR",
        1996: "Dogum",
        2024: "plan1"} //bu sekilde de bir map tanimlamasi yapabiliriz veya sadece {} kullanarak da tanimlama yapilabilir
	map1[0] = "merhaba dunya"
	map2[0] = 23
	map3["nerelerdesin"] = 45
	fmt.Println("map1: ", map1, " map2: ", map2," map3: ",map3, " map4: ",map4," len map1: ", len(map1), " len map2: ", len(map2), " len map3: ", len(map3)) //map'in capacitysi olmaz. 
}