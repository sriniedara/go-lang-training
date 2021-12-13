package mygo

//var count int
//var chan1 chan int

func TestGoChannels() {

	/*
		chan1 = make(chan int)
		go goRoutine1()
		//go goRoutine2()
		for count < 10 {
			<-chan1
			fmt.Printf("%v, ", count)
		}
		//println("Main Routine : count:", count)
	*/
	testBufferedChannels()
}

/*
func goRoutine1() {
	for i := 1; i <= 20; i++ {
		increment()
		chan1 <- count
	}
	println("Go routine1 exit...")
}
*/
/*
func goRoutine2() {

}
*/
/*
func increment() {
	count++
}
*/
//==========================

func testBufferedChannels() {
	var chan1 chan string = make(chan string, 2)
	chan1 <- "srini"
	chan1 <- "Rao"

	go func() {
		chan1 <- "venkat"
		chan1 <- "vijay"
	}()

	println(<-chan1)
	println(<-chan1)
	println(<-chan1)
	println(<-chan1)
	println("********Buffered Channel*********")

}
