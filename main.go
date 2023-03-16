package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var x int16
	x = 8
	z := 5

	println(x)
	println(z)
	value, err := strconv.ParseInt("7", 0, 64)

	if err != nil {
		fmt.Printf("[error]: %v\n", err)
	} else {
		println(value)
	}

	c := make(chan int)
	go doSomething(c)
	<-c

	myMap := make(map[int]bool)
	myMap[1] = true
	println(&myMap)

	s := []int{1, 5, 10}

	for index, value := range s {
		println("index", index)
		println(value)
	}
}

func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	println("Done MID LEVEL 40/100")
	c <- 5
}
