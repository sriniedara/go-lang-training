package mygo

import "fmt"

var count int
var chan1 chan int

func TestGoChannels() {

	chan1 = make(chan int)
	go goRoutine1()
	//go goRoutine2()
	for count < 20 {
		<-chan1
		fmt.Printf("%v, ", count)
	}
	//println("Main Routine : count:", count)
}

func goRoutine1() {
	for i := 1; i <= 20; i++ {
		increment()
		chan1 <- count
	}
}

/*
func goRoutine2() {

}
*/

func increment() {
	count++
}
