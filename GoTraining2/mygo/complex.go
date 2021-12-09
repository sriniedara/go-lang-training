package mygo

import "fmt"

func TestComplex() {
	c1 := 5 + 6i
	fmt.Println("C:", c1)
	c2 := 7 + 8i

	cadd := c1 + c2
	fmt.Println("C-add:", cadd)
	cmul := c1 * c2
	fmt.Println("C-mul:", cmul)
}
