package mygo

import "fmt"

func addMul(a int, b int) (add int, mul int) {
	add = a + b
	mul = a * b

	return add, mul
}

func TestFunctions() {
	add1, mul1 := addMul(5, 6)
	fmt.Println("add:", add1, " mul:", mul1)
	_, mul2 := addMul(8, 9)
	fmt.Println("mul2:", mul2)

	TestAnonymousFunctions()
}
