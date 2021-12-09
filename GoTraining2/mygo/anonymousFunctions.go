package mygo

import "fmt"

func TestAnonymousFunctions() {

	println("Testing Anonymous Functions.....")
	// without argument
	func() {
		fmt.Println("Hello World by anon")
	}()

	// with string as arg

	func(str string) {
		fmt.Println(str)
	}("Hello Folks")

	// anon func with return
	square := func(n int) int {
		return n * n
	}(100)
	fmt.Println("Square", square)

	// anon with multiple returns

	a, p := func(l, b float32) (float32, float32) {
		return l * b, 2 * (l + b)
	}(12.5, 15.5)
	fmt.Println("Area Of Rect and Perimeter of Rect", a, p)

	execute(func() {
		fmt.Println("Hello World,Execute the function with func parameter")
	})

	execute(display)

	add := calc(10, 20, addition)
	fmt.Println("Addition :", add)

	sub := calc(20, 10, func(a, b int) int {
		return a - b
	})

	fmt.Println("Sub:", sub)

	mul := calc(20, 10, func(a, b int) int {
		return a * b
	})

	fmt.Println("Mul:", mul)

	div := calc(20, 10, division)

	fmt.Println("Div:", div)
}

func addition(i, j int) int {
	return i + j
}

func division(a, b int) int {
	return a / b
}

func display() {
	fmt.Println("Calling display function")
}

func execute(f func()) {
	f()
}

func calc(a, b int, f func(int, int) int) int {
	return f(a, b)
}
