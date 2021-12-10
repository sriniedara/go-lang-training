package mygo

import "fmt"

func TestGoRoutines() {
	println("Step-1")
	go goRoutine()
	println("Step-2")
}

func goRoutine() {
	fmt.Println("Hello World")
}
