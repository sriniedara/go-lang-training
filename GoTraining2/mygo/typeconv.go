package mygo

import (
	"fmt"
)

func TestTypeConv() {
	i := 10
	j := 11.12

	k := i + int(j)
	fmt.Printf("type of i:%T", i)
	println("k:", k)

	c := 'A'
	println(c)
	fmt.Printf("c:%c", c)

	//str := "sample"
	//str1 := "123"
	//p := strconv.Atoi(str1)
	//println(p)
}
