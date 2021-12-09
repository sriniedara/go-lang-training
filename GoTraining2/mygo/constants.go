package mygo

import "fmt"

func TestConstants() {

	const i, j int = 10, 20
	const a, b = 10.01, 20.02
	const (
		x      int     = 11
		y      float32 = 22
		name           = "srini"
		height float32 = 5.10
	)
	var result float32 = float32(x) / y
	fmt.Printf("Type=%T \n", result)
	fmt.Println("result=", result)
}
